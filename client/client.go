package client

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"regexp"
	"strings"
)

const ghBaseApiUrl = "https://api.github.com/"

// ERRORS

var (
	ErrInvalidAuthMethod = errors.New("Invalid Github Auth Method")
	ErrNoNextPageUrl = errors.New("No next page url")
	ErrNoPrevPageUrl = errors.New("No prev page url")
	ErrNoFirstPageUrl = errors.New("No first page url")
	ErrNoLastPageUrl = errors.New("No last page url")
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
    authMode          ghAuthModes
}

func NewGithubClient(login string, tokenOrPassword string, authMode ghAuthModes) (ghc *GithubClient, err error) {
	if authMode.IsValid() == false {
		return nil, ErrInvalidAuthMethod
	}

	ghc = &GithubClient{login: login, tokenOrPassword: tokenOrPassword,	authMode: authMode}

	return ghc, nil
}

func (ghc *GithubClient) NewAPIRequest(method, url string, body io.Reader) (req *http.Request, err error) {
	return ghc.newAPIRequest(method, ghBaseApiUrl + url, body)
}

func (ghc *GithubClient) newAPIRequest(method, url string, body io.Reader) (req *http.Request, err error) {
	req, err = http.NewRequest(method, url, body)

	if err != nil {
		return
	}

	switch ; ghc.authMode {
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

	res = newGithubResult(ghc, resp)

	return 
}

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

func (r *GithubResult) parsePageUrls(link string) {
	reUrls, _ := regexp.Compile("<(.*)>; rel=\"(.*)\"")

	urls := strings.Split(link,",")

	for i := range urls {
		m := reUrls.FindStringSubmatch(urls[i])

		if len(m) == 3 {
			switch name := m[2]; name {
			case "first":
				r.firstPageUrl = m[1]
			case "last":
				r.lastPageUrl = m[1]
			case "prev":
				r.prevPageUrl = m[1]
			case "next":
				r.nextPageUrl = m[1]
			}
		} else {
			// TODO
		}
	}
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

func (r *GithubResult) IsPaginated() (ok bool) {
	if (r.nextPageUrl != "" || r.prevPageUrl != "" ||
		r.firstPageUrl != "" || r.lastPageUrl != "") {
		ok = true
    } else {
		ok = false
	}
	return
}

func (r *GithubResult) HasNextPage() (ok bool) {
	if (r.nextPageUrl != "") {
		ok = true
	} else {
		ok = false
	}
	return
}

func (r *GithubResult) HasPrevPage() (ok bool) {
	if (r.prevPageUrl != "") {
		ok = true
	} else {
		ok = false
	}
	return
}

func (r *GithubResult) HasFirstPage() (ok bool) {
	if (r.firstPageUrl != "") {
		ok = true
	} else {
		ok = false
	}
	return
}

func (r *GithubResult) HasLastPage() (ok bool) {
	if (r.lastPageUrl != "") {
		ok = true
	} else {
		ok = false
	}
	return
}

func (r *GithubResult) getPage(fullPageUrl string) (res *GithubResult, err error) {
	req, err := r.ghc.newAPIRequest("GET", fullPageUrl, nil)

	if err != nil {
		return
	}

	httpc := new(http.Client)

	res, err = r.ghc.RunRequest(req, httpc)

	return
}

func (r *GithubResult) GetNextPage() (res *GithubResult, err error) {
	if !r.HasNextPage() {
		return nil, ErrNoNextPageUrl
	}

	return r.getPage(r.nextPageUrl)
}

func (r *GithubResult) GetPrevPage() (res *GithubResult, err error) {
	if !r.HasPrevPage() {
		return nil, ErrNoPrevPageUrl
	}

	return r.getPage(r.prevPageUrl)
}

func (r *GithubResult) GetFirstPage() (res *GithubResult, err error) {
	if !r.HasFirstPage() {
		return nil, ErrNoFirstPageUrl
	}

	return r.getPage(r.firstPageUrl)
}

func (r *GithubResult) GetLastPage() (res *GithubResult, err error) {
	if !r.HasLastPage() {
		return nil, ErrNoLastPageUrl
	}

	return r.getPage(r.lastPageUrl)
}


