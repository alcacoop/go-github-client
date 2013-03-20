// Copyright 2012 Alca Società Cooperativa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gists

import (
	ghclient "github.com/fluffle/go-github-client/client"
	"net/http"
)

// Request the list of authenticated user gists (or public if used anonymously)
// It returns a GithubResult and an error.
func (ghc *Gists) GetGistsList() (res *ghclient.GithubResult, err error) {
	req, err := ghc.NewAPIRequest("GET", "gists", nil)

	if err != nil {
		return
	}

	httpc := new(http.Client)

	res, err = ghc.RunRequest(req, httpc)

	return
}

// Request the list of authenticated user starred gists
// It returns a GithubResult and an error.
func (ghc *Gists) GetStarredGistsList() (res *ghclient.GithubResult, err error) {
	req, err := ghc.NewAPIRequest("GET", "gists/starred", nil)

	if err != nil {
		return
	}

	httpc := new(http.Client)

	res, err = ghc.RunRequest(req, httpc)

	return
}

// Request the list of public gists
// It returns a GithubResult and an error.
func (ghc *Gists) GetPublicGistsList() (res *ghclient.GithubResult, err error) {
	req, err := ghc.NewAPIRequest("GET", "gists/public", nil)

	if err != nil {
		return
	}

	httpc := new(http.Client)

	res, err = ghc.RunRequest(req, httpc)

	return
}

// Request the list of public gists from a given github username.
// It returns a GithubResult and an error.
func (ghc *Gists) GetGistsListByUser(username string) (res *ghclient.GithubResult, err error) {
	req, err := ghc.NewAPIRequest("GET", "users/"+username+"/gists", nil)

	if err != nil {
		return
	}

	httpc := new(http.Client)

	res, err = ghc.RunRequest(req, httpc)

	return
}
