// Copyright 2012 Alca Società Cooperativa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gists

import (
	ghclient "github.com/alcacoop/go-github-client/client"
	"net/http"
	"encoding/json"
	"bytes"
	"errors"
)

var (
	ErrInvalidGistDataCreate = errors.New("Invalid or nil Create Gist Data")
	ErrInvalidGistDataUpdate = errors.New("Invalid or nil Update Gist Data")
)

// Request the content of a gists from a given gist id.
// It returns a GithubResult and an error.
func (ghc *Gists) GetGist(gistId string) (res *ghclient.GithubResult, err error) {
	req, err := ghc.NewAPIRequest("GET", "gists/"+gistId, nil)

	if err != nil {
		return
	}

	httpc := new(http.Client)

	res, err = ghc.RunRequest(req, httpc)

	return
}

// GistDataCreate represents a gists create payload.
// (used as "gists.CreateGist" api call parameter)
type GistDataCreate struct {
	Description string `json:"description,omitempty"` // optional
	Public bool `json:"public"` // required
	Files map[string]GistFileContent `json:"files"` // required
}

// create a new GistDataCreate struct
func NewGistDataCreate() *GistDataCreate {
	data := new(GistDataCreate)
	data.Files = make(map[string]GistFileContent)

	return data
}

// validate gists create payload
func (cgd *GistDataCreate) IsValid() bool {
	// TODO: add validate data
	return true
}

// GistDataCreate represents a gists update payload
// (used as "gists.UpdateGist" api call parameter)
type GistDataUpdate struct {
	Description string `json:"description,omitempty"` // optional
	Files map[string]GistFileContent `json:"files"` // required
}

// create a new GistDataUpdate struct
func NewGistDataUpdate() *GistDataUpdate {
	data := new(GistDataUpdate)
	data.Files = make(map[string]GistFileContent)

	return data
}

// validate gists update payload
func (cgd *GistDataUpdate) IsValid() bool {
	// TODO: add validate data
	return true
}

// GistFileContent represents a create/update gists files payload. 
// It should be added to GistsDataCreate or GistsDataUpdate objects.
// 
// As example to create data needed to create a new gist:
//   newGist := ghgists.NewGistDataCreate()
//   newGist.Description = "test go create"
//   newGist.Public = false
//   newGist.Files["prova1.js"] = ghgists.GistFileContent{Content: "var x=5;"}
// 
// and to create data needed to update an existent gist:
//   updGist := ghgists.NewGistDataUpdate()
//   updGist.Description = "test go update"
//   updGist.Files["prova1.js"] = ghgists.GistFileContent{Filename: "renamed.js",
//                                                        Content: "var x=5;"}
type GistFileContent struct {
	Filename string `json:"filename,omitempty"` // optional (used on UpdateGistData)
	Content string `json:"content"` // required
}

// create a new gists using a given GistDataCreate struct.
// It returns a GithubResult and an error.
func (ghc *Gists) CreateGist(data *GistDataCreate) (res *ghclient.GithubResult, err error) {
	if data == nil || !data.IsValid() {
		err = ErrInvalidGistDataCreate
		return
	}

	body, err := json.Marshal(data)
	
	if err != nil {
		return
	}

	req, err := ghc.NewAPIRequest("POST", "gists", bytes.NewBuffer(body))

	if err != nil {
		return
	}

	httpc := new(http.Client)

	res, err = ghc.RunRequest(req, httpc)

	return	
}

// update an existent gists using a given gist id and  GistDataCreate struct.
// It returns a GithubResult and an error.
func (ghc *Gists) UpdateGist(gistid string, data *GistDataUpdate) (res *ghclient.GithubResult, err error) {
	if data == nil || !data.IsValid() {
		err = ErrInvalidGistDataUpdate
		return
	}

	body, err := json.Marshal(data)
	
	if err != nil {
		return
	}

	req, err := ghc.NewAPIRequest("POST", "gists/"+gistid, bytes.NewBuffer(body))

	if err != nil {
		return
	}

	httpc := new(http.Client)

	res, err = ghc.RunRequest(req, httpc)

	return	
}