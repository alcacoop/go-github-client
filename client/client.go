// Copyright 2012 Alca Società Cooperativa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package client

import (
	"errors"
	"io"
	"net/http"
	"fmt"
)

const ghBaseApiUrl = "https://api.github.com/"

// ERRORS

var (
	ErrInvalidAuthMethod = errors.New("Invalid Github Auth Method")
	ErrNoNextPageUrl     = errors.New("No next page url")
	ErrNoPrevPageUrl     = errors.New("No prev page url")
	ErrNoFirstPageUrl    = errors.New("No first page url")
	ErrNoLastPageUrl     = errors.New("No last page url")
)

// AUTH METHODS
type ghAuthModes int

const (
	AUTH_OAUTH2_TOKEN = iota
	AUTH_USER_PASSWORD
)

func (s ghAuthModes) IsValid() (ok bool) {
	if s >= AUTH_OAUTH2_TOKEN && s <= AUTH_USER_PASSWORD {
		ok = true
	} else {
		ok = false
	}

	return
}

type GithubClient struct {
	login           string
	tokenOrPassword string
	authMode        ghAuthModes
}

func NewGithubClient(login string, tokenOrPassword string, authMode ghAuthModes) (ghc *GithubClient, err error) {
	if authMode.IsValid() == false {
		return nil, ErrInvalidAuthMethod
	}

	ghc = &GithubClient{login: login, tokenOrPassword: tokenOrPassword, authMode: authMode}

	return ghc, nil
}

func (ghc *GithubClient) NewAPIRequest(method, url string, body io.Reader) (req *http.Request, err error) {
	return ghc.newAPIRequest(method, ghBaseApiUrl+url, body)
}

func (ghc *GithubClient) newAPIRequest(method, url string, body io.Reader) (req *http.Request, err error) {
	req, err = http.NewRequest(method, url, body)

	if err != nil {
		return
	}

	switch ghc.authMode {
	case AUTH_OAUTH2_TOKEN:
		req.Header.Add("Authorization", "token "+ghc.tokenOrPassword)
	case AUTH_USER_PASSWORD:
		req.SetBasicAuth(ghc.login, ghc.tokenOrPassword)
	}

	return
}

func (ghc *GithubClient) RunRequest(req *http.Request, httpc *http.Client) (res *GithubResult, err error) {
	resp, err := httpc.Do(req)

	if err != nil {
		return
	}

	fmt.Printf("REQ %+v\n\n",req)
	fmt.Printf("RESP %+v\n\n",resp)
	res = newGithubResult(ghc, resp)

	return
}

// TBD
func (ghc *GithubClient) GetResourceFromUrl(fullResourceUrl string) (res *GithubResult, err error) {
	// TODO: validate resource url
	req, err := ghc.newAPIRequest("GET", fullResourceUrl, nil)

	if err != nil {
		return
	}

	httpc := new(http.Client)

	res, err = ghc.RunRequest(req, httpc)

	return
}
