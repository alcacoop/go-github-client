package users

import (
	ghclient "github.com/alcacoop/go-github-client/client"
	"net/http"
)

type Users struct {
  *ghclient.GithubClient
}

func NewUsers(ghc *ghclient.GithubClient) *Users {
	userc := new(Users)
	userc.GithubClient = ghc

	return userc
}

func (ghc *Users) GetAuthenticatedUserInfo() (res *ghclient.GithubResult, err error) {
	req, err := ghc.NewAPIRequest("GET", "user", nil)

	if err != nil {
		return
	}

	httpc := new(http.Client)

	res, err = ghc.RunRequest(req, httpc)

	return
}

func (ghc *Users) GetUserInfo(username string) (res *ghclient.GithubResult, err error) {
	req, err := ghc.NewAPIRequest("GET", "users/"+username, nil)

	if err != nil {
		return
	}

	httpc := new(http.Client)

	res, err = ghc.RunRequest(req, httpc)

	return
}

// TODO: patch user info
/*
Update the authenticated user

PATCH /user

Input

name
    Optional string 
email
    Optional string - Publicly visible email address.
blog
    Optional string 
company
    Optional string 
location
    Optional string 
hireable
    Optional boolean 
bio
    Optional string 

{
  "name": "monalisa octocat",
  "email": "octocat@github.com",
  "blog": "https://github.com/blog",
  "company": "GitHub",
  "location": "San Francisco",
  "hireable": true,
  "bio": "There once..."
}
*/