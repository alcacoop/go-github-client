// Copyright 2012 Alca Società Cooperativa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package users

import (
	ghclient "github.com/alcacoop/go-github-client/client"
	"net/http"
)

type PlanDetail struct {
	Name          string
	Space         int
	Collaborators int
	Private_repos int
}

type UserShort struct {
	Login       string
	Id          int
	Avatar_url  string
	Gravatar_id string
	Url         string
}

type UserDetail struct {
	Id                  int
	Login               string
	Avatar_url          string
	Gravatar_id         string
	Url                 string
	Received_events_url string
	Type                string
	Following           int
	Site_admin          bool
	Followers_url       string
	Created_at          string
	Plan                *PlanDetail
	Following_url       string
	Organizations_url   string
	Events_url          string
	Updated_at          string
	Private_gists       int
	Gists_url           string
	Starred_url         string
	Subscriptions_url   string
	Disk_usage          int
	Html_url            string
	Repos_url           string
	Followers           int
	Public_gists        int
	Total_private_repos int
	Owned_private_repos int
	Collaborators       int
	Public_repos        int
	Fork                bool
}

// Users is a simplified github users api client.
type Users struct {
	*ghclient.GithubClient
}

// create a new github Users client from an existent GithubClient
func NewUsers(ghc *ghclient.GithubClient) *Users {
	userc := new(Users)
	userc.GithubClient = ghc
	return userc
}

// Request the current autenticated user info.
// It returns a UserDetail and an error.
func (ghc *Users) GetAuthenticatedUserInfo() (*UserDetail, error) {
	req, err := ghc.NewAPIRequest("GET", "user", nil)

	if err != nil {
		return nil, err
	}

	httpc := new(http.Client)
	res, err := ghc.RunRequest(req, httpc)

	var detail UserDetail
	_, err = res.JsonStruct(&detail)
	return &detail, err
}

// Request github user info about a defined username.
// It returns a UserDetail and an error.
func (ghc *Users) GetUserInfo(username string) (*UserDetail, error) {
	req, err := ghc.NewAPIRequest("GET", "users/"+username, nil)

	if err != nil {
		return nil, err
	}

	httpc := new(http.Client)

	res, err := ghc.RunRequest(req, httpc)

	var detail UserDetail
	_, err = res.JsonStruct(&detail)
	return &detail, err
}
