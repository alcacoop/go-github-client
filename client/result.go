package client

import (
	"io/ioutil"
	"net/http"
	"encoding/json"
)

type GithubResult struct {
	ghc *GithubClient
	RawHttpResponse *http.Response

	jsonBody interface{}
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

func (r *GithubResult) Json() (jr interface{}, err error) {
	if r.jsonBody == nil && r.jsonParseError == nil {
		err = r.parseBody()

		if err != nil {
			r.jsonParseError = err
			return
		}
	}

	return r.jsonBody, r.jsonParseError
}
