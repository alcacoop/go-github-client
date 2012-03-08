// Copyright 2012 Alca Società Cooperativa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gists

import (
	ghclient "github.com/alcacoop/go-github-client/client"
)

// Gists is a simplified github gists api client
type Gists struct {
	*ghclient.GithubClient
}

// create a new github Gists client from an existent GithubClient
func NewGists(ghc *ghclient.GithubClient) *Gists {
	gists := new(Gists)
	gists.GithubClient = ghc

	return gists
}
