// Copyright 2012 Alca Società Cooperativa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package issues

import (
	ghclient "github.com/alcacoop/go-github-client/client"
)

// Issues is a simplified github issues api client
type Issues struct {
  *ghclient.GithubClient
}

// create a new github issues client from an existent GithubClient
func NewIssues(ghc *ghclient.GithubClient) *Issues {
	issues := new(Issues)
	issues.GithubClient = ghc

	return issues
}
