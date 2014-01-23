// Copyright 2012 Alca Società Cooperativa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package repos

import (
	ghclient "github.com/iansmith/go-github-client/client"
	ghuser "github.com/iansmith/go-github-client/users"
	"fmt"
	"net/http"
)


type RepoDetail struct {
	Id            int
	Url           string
	Html_url      string
	Clone_url     string
	Git_url       string
	Ssh_url       string
	Svn_url       string
	Mirror_url    string
	Parent        *RepoDetail
	Source        *RepoDetail
	Owner         *ghuser.UserShort
	Name          string
	Full_name     string
	Description   string
	Homepage      string
	Language      string
	Private       bool
	Fork          bool
	Forks         int
	Hireable      bool
	Watchers      int
	Size          int
	Master_branch string
	Open_issues   int
	Pushed_at     string
	Created_at    string
	Updated_at    string
}

// Repos is a simplified github repos api client.
type Repos struct {
	*ghclient.GithubClient
	login string
}

// create a new github Repos client from an existent GithubClient
func NewRepos(ghc *ghclient.GithubClient) *Repos {
	repoc := new(Repos)
	repoc.GithubClient = ghc
	return repoc
}

func (self *Repos) readLogin() (error) {
	if self.login!="" {
		return nil
	}
	users := ghuser.NewUsers(self.GithubClient)
	userDetail, err := users.GetAuthenticatedUserInfo()
	if err!=nil {
		return err
	}
	self.login = userDetail.Login
	return nil
}

// Request the current autenticated user's repos.
// It returns a RepoDetail, the ghclient.Result (for pagination info) and an error.
func (self *Repos) GetAuthenticatedUserRepos() ([]*RepoDetail, *ghclient.GithubResult, error) {

	req, err := self.NewAPIRequest("GET", "user/repos", nil)
	if err != nil {
		return nil, nil, err
	}
	httpc := new(http.Client)
	res, err := self.RunRequest(req, httpc)

	detail := []*RepoDetail{}
	_, err = res.JsonStruct(&detail)
	return detail, res, err

}

func (self *Repos) GetAuthenticatedUserPublicInfo(name string) (*RepoDetail, error) {
	err := self.readLogin()
	if err != nil {
		return nil, err
	}
	req, err := self.NewAPIRequest("GET", fmt.Sprintf("repos/%s/%s", self.login, name), nil)
	if err != nil {
		return nil, err
	}
	httpc := new(http.Client)
	res, err := self.RunRequest(req, httpc)

	var detail RepoDetail
	_, err = res.JsonStruct(&detail)
	return &detail, err
}
