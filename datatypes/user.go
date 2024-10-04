package datatypes

import (
	"fmt"
	"strings"
	"time"
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
