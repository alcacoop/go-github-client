// Copyright 2012 Alca Società Cooperativa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package issues

import (
	ghclient "github.com/alcacoop/go-github-client/client"
	"net/http"
	"encoding/json"
	"bytes"
	"errors"
)

var ErrInvalidIssuesListOptions error = errors.New("Invalid Issue List Options")

type UserListFilter string

var (
	FILTER_ASSIGNED UserListFilter = "assigned"
	FILTER_CREATED UserListFilter = "created"
	FILTER_MENTIONED UserListFilter = "mentioned"
	FILTER_SUBSCRIBED UserListFilter = "subscribed"
)

type IssueState string

var (
	STATE_OPEN IssueState = "open"
	STATE_CLOSED IssueState = "closed"
)

type ListSort string

var (
	SORT_CREATED ListSort = "created"
	SORT_UPDATED ListSort = "updated"
	SORT_COMMENTS ListSort = "comments"
)

type ListDirection string

var (
	DIRECTION_ASC ListDirection = "asc"
	DIRECTION_DESC ListDirection = "desc"
)

// TBD
type UserListOptions struct {
	Filter UserListFilter `json:"filter,omitempty"` // assigned, created, mentioned, subscribed	
	State  IssueState `json:"state,omitempty"` // open, closed
	Labels string     `json:"labels,omitempty"` // comma separated labels
	Sort   ListSort   `json:"sort,omitempty"` // created, updated, comments, default: created.
	Direction   ListSort   `json:"direction,omitempty"` // asc, desc, default: desc
	Since  string     `json:"since,omitempty"` // timestamp in ISO 8601 format: YYYY-MM-DDTHH:MM:SSZ
}

// TBD
func (s *UserListOptions) IsValid() bool {
	// TODO
	return true
}

// TBD
func (ghc *Issues) GetUserIssuesList(options *UserListOptions) (res *ghclient.GithubResult, err error) {
	if options == nil || !options.IsValid() {
		err = ErrInvalidIssuesListOptions
		return
	}

	body, err := json.Marshal(options)
	
	if err != nil {
		return
	}

	req, err := ghc.NewAPIRequest("GET", "issues", bytes.NewBuffer(body))

	//req.Header.Add("Accept", "*/*")
	//req.Header.Add("Content-Type", "application/json")

	if err != nil {
		return
	}

	httpc := new(http.Client)

	res, err = ghc.RunRequest(req, httpc)

	return
}

/////////////////////
// REPO ISSUE LIST
/////////////////////

var (
	ALL string = "*"
	NONE string = "none"
)

type RepoListOptions struct {
	Milestone string `json:"milestone,omitempty"` // MILESTONE_ID, none, *
	State  IssueState `json:"state,omitempty"` // open, closed
	Assegnee string `json:"assegnee,omitempty"` // USERNAME, none, *
	Mentioned string `json:"mentioned,omitempty"` // USERNAME
	Labels string     `json:"labels,omitempty"` // comma separated labels
	Sort   ListSort   `json:"sort,omitempty"` // created, updated, comments, default: created.
	Direction   ListSort   `json:"direction,omitempty"` // asc, desc, default: desc
	Since  string     `json:"since,omitempty"` // timestamp in ISO 8601 format: YYYY-MM-DDTHH:MM:SSZ
}

func (s *RepoListOptions) IsValid() bool {
	// TODO
	return true
}

// TBD
func (ghc *Issues) GetRepoIssuesList(user string, repo string, 
	options *RepoListOptions) (res *ghclient.GithubResult, err error) {
	if options == nil || !options.IsValid() {
		err = ErrInvalidIssuesListOptions
		return
	}

	body, err := json.Marshal(options)
	
	if err != nil {
		return
	}

	req, err := ghc.NewAPIRequest("GET", "repos/"+user+"/"+repo+
		"/issues", bytes.NewBuffer(body))

	//req.Header.Add("Accept", "*/*")
	//req.Header.Add("Content-Type", "application/json")

	if err != nil {
		return
	}

	httpc := new(http.Client)

	res, err = ghc.RunRequest(req, httpc)

	return
}
