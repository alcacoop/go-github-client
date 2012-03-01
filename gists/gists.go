package gists

import (
	ghclient "github.com/alcacoop/go-github-client/client"
	"net/http"
)

type Gists struct {
  *ghclient.GithubClient
}

func NewGists(ghc *ghclient.GithubClient) *Gists {
	gists := new(Gists)
	gists.GithubClient = ghc

	return gists
}

func (ghc *Gists) GetGistsList() (res *ghclient.GithubResult, err error) {
	req, err := ghc.NewAPIRequest("GET", "gists", nil)

	if err != nil {
		return
	}

	httpc := new(http.Client)

	res, err = ghc.RunRequest(req, httpc)

	return
}

func (ghc *Gists) GetStarredGistsList() (res *ghclient.GithubResult, err error) {
	req, err := ghc.NewAPIRequest("GET", "gists/starred", nil)

	if err != nil {
		return
	}

	httpc := new(http.Client)

	res, err = ghc.RunRequest(req, httpc)

	return
}

func (ghc *Gists) GetPublicGistsList() (res *ghclient.GithubResult, err error) {
	req, err := ghc.NewAPIRequest("GET", "gists/public", nil)

	if err != nil {
		return
	}

	httpc := new(http.Client)

	res, err = ghc.RunRequest(req, httpc)

	return
}

func (ghc *Gists) GetGistsListByUser(username string) (res *ghclient.GithubResult, err error) {
	req, err := ghc.NewAPIRequest("GET", "users/"+username+"/gists", nil)

	if err != nil {
		return
	}

	httpc := new(http.Client)

	res, err = ghc.RunRequest(req, httpc)

	return
}

func (ghc *Gists) GetGist(gistId string) (res *ghclient.GithubResult, err error) {
	req, err := ghc.NewAPIRequest("GET", "gists/"+gistId, nil)

	if err != nil {
		return
	}

	httpc := new(http.Client)

	res, err = ghc.RunRequest(req, httpc)

	return
}