// Copyright 2012 Alca Società Cooperativa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package client

import (
	"encoding/base64"
	"testing"
)

func TestGhAuthModesValidation(t *testing.T) {
	_, err := NewGithubClient("testuser", "password", AUTH_USER_PASSWORD)

	if err != nil {
		t.Error("It should return no error")
	}

	_, err = NewGithubClient("testuser", "password", 5)

	if err == nil {
		t.Error("It should return an error")
	}

	if err != ErrInvalidAuthMethod {
		t.Error("It should return an ErrInvalidAuthMethod")
	}
}

func TestCreateNewAPIRequest(t *testing.T) {
	// TEST BASIC AUTH
	ghc, _ := NewGithubClient("testuser", "password", AUTH_USER_PASSWORD)

	req, err := ghc.NewAPIRequest("GET", "http://localhost", nil)

	if err != nil {
		t.Error("CreateNewAPIRequest should not return an error")
	}

	if req.Header.Get("Authorization") != "Basic "+
		base64.StdEncoding.EncodeToString([]byte("testuser:password")) {
		t.Error("Request should contain correct basic auth info")
	}

	// TEST OAUTH2
	ghc, _ = NewGithubClient("testuser", "MYTOKEN", AUTH_OAUTH2_TOKEN)

	req, err = ghc.NewAPIRequest("GET", "http://localhost", nil)

	if err != nil {
		t.Error("CreateNewAPIRequest should not return an error")
	}

	if req.Header.Get("Authorization") != "token MYTOKEN" {
		t.Error("Request should not contains correct basic auth info")
	}
}
