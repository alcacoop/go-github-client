package client

import (
	"testing"
	"encoding/base64"
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

	req, err := ghc.CreateNewAPIRequest("GET", "http://localhost", nil)

	if err != nil {
		t.Error("CreateNewAPIRequest should not return an error")
	}

	if req.Header.Get("Authorization") != "Basic "+ 
		base64.StdEncoding.EncodeToString([]byte("testuser:password")) {
		t.Error("Request should contain correct basic auth info")
	}

	// TEST OAUTH2
	ghc, _ = NewGithubClient("testuser", "MYTOKEN", AUTH_OAUTH2_TOKEN)

	req, err = ghc.CreateNewAPIRequest("GET", "http://localhost", nil)

	if err != nil {
		t.Error("CreateNewAPIRequest should not return an error")
	}

	if req.Header.Get("Authorization") != "token MYTOKEN" {
		t.Error("Request should not contains correct basic auth info")
	}
}