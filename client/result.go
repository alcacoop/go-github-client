// Copyright 2012 Alca Società Cooperativa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package client

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
)

/* 
GithubResult abstracts a generic reply from the Github REST API, parse it
internally and adds some helper (e.g. json data, paginated results, rate limits)

GithubResult objects are returned by GithubClient.RunRequest method.
*/
type GithubResult struct {
	ghc *GithubClient
	RawHttpResponse *http.Response

	jsonBody JsonData
	jsonParseError error

	RateLimitLimit string
	RateLimitRemaining string

	firstPageUrl string
	lastPageUrl string
	prevPageUrl string
    nextPageUrl string
}

func newGithubResult(ghc *GithubClient, resp *http.Response) *GithubResult {
	result := new(GithubResult)
	result.ghc = ghc
	result.RawHttpResponse = resp	

	result.parseHeader()
	result.parseBody()

	return result
}

// tests if server response status is 200
func (r *GithubResult) IsSuccess() bool {
	return r.RawHttpResponse.StatusCode == 200
}

func (r *GithubResult) parseBody() (err error) {
	data, err := ioutil.ReadAll(r.RawHttpResponse.Body)

	if err != nil {
		return
	}

	err = json.Unmarshal(data, &r.jsonBody)

	return
}

func (r *GithubResult) parseHeader() {
	r.RateLimitLimit = r.RawHttpResponse.Header.Get("X-Ratelimit-Limit")
	r.RateLimitRemaining = r.RawHttpResponse.Header.Get("X-Ratelimit-Remaining")
	
	r.parsePageUrls(r.RawHttpResponse.Header.Get("Link"))

    return
}

// Lazily parse body as json and return unmarshalled results.
func (r *GithubResult) Json() (jr JsonData, err error) {
	if r.jsonBody == nil && r.jsonParseError == nil {
		err = r.parseBody()

		if err != nil {
			r.jsonParseError = err
			return
		}
	}

	return r.jsonBody, r.jsonParseError
}
