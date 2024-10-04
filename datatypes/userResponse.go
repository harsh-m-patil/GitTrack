package datatypes

import "time"

type UserResponse struct {
	Login           string    `json:"login"`
	Type            string    `json:"type"`
	Name            string    `json:"name"`
	Blog            string    `json:"blog"`
	Location        string    `json:"location"`
	Email           string    `json:"email"`
	Bio             string    `json:"bio"`
	TwitterUsername string    `json:"twitter_username"`
	PublicRepos     int       `json:"public_repos"`
	PublicGists     int       `json:"public_gists"`
	Followers       int       `json:"followers"`
	Following       int       `json:"following"`
	CreatedAt       time.Time `json:"created_at"`
}
