package gists

import (
	ghclient "github.com/alcacoop/go-github-client/client"
)

type Gists struct {
  *ghclient.GithubClient
}

func NewGists(ghc *ghclient.GithubClient) *Gists {
	gists := new(Gists)
	gists.GithubClient = ghc

	return gists
}
