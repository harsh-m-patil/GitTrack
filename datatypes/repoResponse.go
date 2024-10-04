package datatypes

import "time"

type RepoResponse []struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Owner    struct {
		Login string `json:"login"`
		ID    int    `json:"id"`
	} `json:"owner"`
	Description     string    `json:"description"`
	Fork            bool      `json:"fork"`
	URL             string    `json:"url"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	PushedAt        time.Time `json:"pushed_at"`
	Size            int       `json:"size"`
	StargazersCount int       `json:"stargazers_count"`
	WatchersCount   int       `json:"watchers_count"`
	Language        string    `json:"language"`
	HasIssues       bool      `json:"has_issues"`
	HasProjects     bool      `json:"has_projects"`
	HasDownloads    bool      `json:"has_downloads"`
	ForksCount      int       `json:"forks_count"`
	Archived        bool      `json:"archived"`
	Disabled        bool      `json:"disabled"`
	OpenIssuesCount int       `json:"open_issues_count"`
	AllowForking    bool      `json:"allow_forking"`
	Topics          []string  `json:"topics"`
	Forks           int       `json:"forks"`
	OpenIssues      int       `json:"open_issues"`
	Watchers        int       `json:"watchers"`
	DefaultBranch   string    `json:"default_branch"`
}
