package datatypes

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type User struct {
	UserName        string
	Name            string
	Email           string
	Bio             string
	PublicRepos     int
	PublicGists     int
	Followers       int
	Following       int
	CreatedAt       time.Time
	TotalCommits    int
	TotalStargazers int
	UsedLanguages   []string
}

func NewUser() *User {
	return &User{}
}

func (user *User) SetUser(userResp UserResponse, repoResp RepoResponse) {
	user.UserName = userResp.Login
	user.Name = userResp.Name
	user.Email = userResp.Email
	user.Bio = userResp.Bio
	user.Followers = userResp.Followers
	user.Following = userResp.Following
	user.CreatedAt = userResp.CreatedAt
	user.PublicRepos = userResp.PublicRepos
	user.PublicGists = userResp.PublicGists
	// TODO:
	user.TotalCommits = 0

	star_counts := 0

	languageSet := newSet[string]()
	for _, repo := range repoResp {
		star_counts += repo.StargazersCount

		if languageSet.Contains(repo.Language) {
			continue
		}

		languageSet.Add(repo.Language)
	}

	languages := languageSet.ToSlice()

	user.TotalStargazers = star_counts
	user.UsedLanguages = languages
}

func (user *User) ToString() string {
	// Format the creation date
	createdAtStr := user.CreatedAt.Format("2006-01-02 15:04:05")

	// Format the languages as a comma-separated string
	languages := strings.Join(user.UsedLanguages, ", ")

	// Return the formatted string
	return fmt.Sprintf(`
    username: %s
    Name: %s
    Email: %s
    Bio: %s
    Public Repos: %d
    Public Gists: %d
    Followers: %d
    Following: %d
    Created At: %s
    Total Commits: %d
    Total Stargazers: %d
    Used Languages: %s`,
		user.UserName, user.Name, user.Email, user.Bio, user.PublicRepos,
		user.PublicGists, user.Followers, user.Following, createdAtStr,
		user.TotalCommits, user.TotalStargazers, languages)
}

func (user *User) PrintWithStyle() {
	// Style for the username
	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#24292E")). // Darker text for better contrast
		Background(lipgloss.Color("#F6F8FA")). // Light background color
		PaddingTop(1).
		PaddingLeft(20).
		Width(53)

	fmt.Println(style.Render(user.UserName))

	re := lipgloss.NewRenderer(os.Stdout)

	const (
		darkBlue = lipgloss.Color("36")  // Dark blue for header
		gray     = lipgloss.Color("245") // Gray for table content
		black    = lipgloss.Color("0")   // Black for general text
	)

	var (
		HeaderStyle = re.NewStyle().Foreground(darkBlue).Bold(true).Align(lipgloss.Center)
		CellStyle   = re.NewStyle().Padding(0, 1).Width(25).Foreground(black) // Black text for better visibility
		RowStyle    = CellStyle.Foreground(gray)                              // Rows styled with gray
		BorderStyle = lipgloss.NewStyle().Foreground(darkBlue)                // Border matching header
	)

	rows := [][]string{
		{"Name", user.Name},
		{"Email", user.Email},
		{"Bio", user.Bio},
		{"Public Repos", fmt.Sprintf("%d", user.PublicRepos)},
		{"Public Gists", fmt.Sprintf("%d", user.PublicGists)},
		{"Followers", fmt.Sprintf("%d", user.Followers)},
		{"Following", fmt.Sprintf("%d", user.Following)},
		{"Created At", user.CreatedAt.Format("2006-01-02 15:04:05")},
		{"Total Commits", fmt.Sprintf("%d", user.TotalCommits)},
		{"Total Stargazers", fmt.Sprintf("%d", user.TotalStargazers)},
		{"Used Languages", strings.Join(user.UsedLanguages, ", ")}, // Corrected to join languages
	}

	t := table.New().
		Border(lipgloss.ThickBorder()).
		BorderStyle(BorderStyle).
		StyleFunc(func(row, col int) lipgloss.Style {
			var style lipgloss.Style

			switch {
			case row == 0:
				return HeaderStyle
			default:
				style = RowStyle
			}

			return style
		}).
		Headers("KEY", "VALUE").
		Rows(rows...)

	fmt.Println(t)
}
