// Copyright 2012 Alca Società Cooperativa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

/* 
GithubResult abstracts a generic reply from the Github REST API, parse it
internally and adds some helper (e.g. json data, paginated results, rate limits)

GithubResult objects are returned by GithubClient.RunRequest method.
*/
type GithubResult struct {
	ghc             *GithubClient
	RawHttpResponse *http.Response

	jsonBody       interface{}
	jsonParseError error

	RateLimitLimit     string
	RateLimitRemaining string

	firstPageUrl string
	lastPageUrl  string
	prevPageUrl  string
	nextPageUrl  string
}

func newGithubResult(ghc *GithubClient, resp *http.Response) *GithubResult {
	result := new(GithubResult)
	result.ghc = ghc
	result.RawHttpResponse = resp

	result.parseHeader()
	//result.parseBody()

	return result
}

// tests if server response status is 200
func (r *GithubResult) IsSuccess() bool {
	return r.RawHttpResponse.StatusCode == 200
}

func (r *GithubResult) parseBody() (err error) {
	if r.RawHttpResponse.ContentLength == 0 {
		// EMPTY OBJECT
		err = json.Unmarshal(([]byte)("[]"), &(r.jsonBody))
	} else if r.RawHttpResponse.ContentLength != 0 {
		defer r.RawHttpResponse.Body.Close()
		data, err := ioutil.ReadAll(r.RawHttpResponse.Body)

		if err != nil {
			return err
		}

		r.jsonParseError = json.Unmarshal(data, &(r.jsonBody))
		err = r.jsonParseError
	}

	return err
}

func (r *GithubResult) parseHeader() {
	r.RateLimitLimit = r.RawHttpResponse.Header.Get("X-Ratelimit-Limit")
	r.RateLimitRemaining = r.RawHttpResponse.Header.Get("X-Ratelimit-Remaining")

	r.parsePageUrls(r.RawHttpResponse.Header.Get("Link"))

	return
}

// Lazily parse body as json and return unmarshalled results.
func (r *GithubResult) Json() (j interface{}, err error) {
	if r.jsonBody == nil && r.jsonParseError == nil {
		err = r.parseBody()

		if err != nil {
			r.jsonParseError = err
			return
		}
	}

	return r.jsonBody, r.jsonParseError
}

// TBD
func (r *GithubResult) JsonMap() (j JsonMap, err error) {
	res, err := r.Json()

	return (JsonMap)(res.(map[string]interface{})), err
}

// TBD
func (r *GithubResult) JsonArray() (j JsonArray, err error) {
	res, err := r.Json()

	return (JsonArray)(res.([]interface{})), err
}
