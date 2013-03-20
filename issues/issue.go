// Copyright 2012 Alca Società Cooperativa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package issues

import (
	"bytes"
	"encoding/json"
	"errors"
	ghclient "github.com/alcacoop/go-github-client/client"
	"net/http"
)

var (
	ErrInvalidIssueDataCreate = errors.New("Invalid or nil Create Issue Data")
	ErrInvalidIssueDataUpdate = errors.New("Invalid or nil Update Issue Data")
)

// Request the content of a issue from a given issue id.
// It returns a GithubResult and an error.
func (ghc *Issues) GetIssue(user, repo, issueId string) (res *ghclient.GithubResult, err error) {
	req, err := ghc.NewAPIRequest("GET", "repos/"+user+"/"+repo+"/issues/"+
		issueId, nil)

	if err != nil {
		return
	}

	httpc := new(http.Client)

	res, err = ghc.RunRequest(req, httpc)

	return
}

// TBD
type IssueDataCreate struct {
	Title     string   `json:"title"`           // required
	Body      string   `json:"body,omitempty"`  // optional
	Assignee  string   `json:"assignee,omitempty"`  // optional
	Milestone int      `json:"milestone,omitempty"`  // optional
	Labels    []string `json:"labels,omitempty"` // optional
}

// TBD
func (idc *IssueDataCreate) IsValid() bool {
	// TODO: add validate data
	return true
}

// TBD
func (ghc *Issues) CreateIssue(user, repo string, data *IssueDataCreate) (res *ghclient.GithubResult, err error) {
	if data == nil || !data.IsValid() {
		err = ErrInvalidIssueDataCreate
		return
	}

	body, err := json.Marshal(data)

	if err != nil {
		return
	}

	req, err := ghc.NewAPIRequest("POST", "repos/"+user+"/"+repo+"/issues",
		bytes.NewBuffer(body))

	if err != nil {
		return
	}

	httpc := new(http.Client)

	res, err = ghc.RunRequest(req, httpc)

	return
}

/////////////////
// UPDATE ISSUE
/////////////////

// TBD
type IssueDataUpdate struct {
	Title     string     `json:"title,omitempty"` // optional
	Body      string     `json:"body,omitempty"`  // optional
	Assignee  string     `json:"assignee,omitempty"`  // optional
	State     IssueState `json:"state,omitempty"` // optional
	Milestone int        `json:"milestone,omitempty"`  // optional
	Labels    []string   `json:"labels,omitempty"` // optional
}

// TBD
func (idc *IssueDataUpdate) IsValid() bool {
	// TODO: add validate data
	return true
}

// TBD
func (ghc *Issues) UpdateIssue(user, repo, issueId string, data *IssueDataUpdate) (res *ghclient.GithubResult, err error) {
	if data == nil || !data.IsValid() {
		err = ErrInvalidIssueDataUpdate
		return
	}

	body, err := json.Marshal(data)

	if err != nil {
		return
	}

	req, err := ghc.NewAPIRequest("PATCH", "repos/"+user+"/"+repo+"/issues"+issueId,
		bytes.NewBuffer(body))

	if err != nil {
		return
	}

	httpc := new(http.Client)

	res, err = ghc.RunRequest(req, httpc)

	return
}
