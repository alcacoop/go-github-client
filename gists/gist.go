package gists

import (
	ghclient "github.com/alcacoop/go-github-client/client"
	"net/http"
	"encoding/json"
	"bytes"
	"errors"
)

var (
	ErrInvalidCreateGistData = errors.New("Invalid o nil Create Gist Data")
)

func (ghc *Gists) GetGist(gistId string) (res *ghclient.GithubResult, err error) {
	req, err := ghc.NewAPIRequest("GET", "gists/"+gistId, nil)

	if err != nil {
		return
	}

	httpc := new(http.Client)

	res, err = ghc.RunRequest(req, httpc)

	return
}

type CreateGistData struct {
	Description string `json:"description,omitempty"` // optional
	Public bool `json:"public"` // required
	Files map[string]GistFileContent `json:"files"` // required
}

func (cgd *CreateGistData) IsValid() bool {
	return true
}

type GistFileContent struct {
	Content string `json:"content"` // required
}

func (ghc *Gists) CreateGist(data *CreateGistData) (res *ghclient.GithubResult, err error) {
	if data == nil || !data.IsValid() {
		err = ErrInvalidCreateGistData
		return
	}

	body, err := json.Marshal(data)
	
	if err != nil {
		return
	}

	req, err := ghc.NewAPIRequest("POST", "gists/", bytes.NewBuffer(body))

	if err != nil {
		return
	}

	httpc := new(http.Client)

	res, err = ghc.RunRequest(req, httpc)

	return	
}