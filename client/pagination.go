// Copyright 2012 Alca Società Cooperativa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package client

import (
	"net/http"
	"regexp"
	"strings"
)

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

// returns true if the result is paginated
func (r *GithubResult) IsPaginated() (ok bool) {
	if (r.nextPageUrl != "" || r.prevPageUrl != "" ||
		r.firstPageUrl != "" || r.lastPageUrl != "") {
		ok = true
    } else {
		ok = false
	}
	return
}

// returns true if the result has a next page
func (r *GithubResult) HasNextPage() (ok bool) {
	if (r.nextPageUrl != "") {
		ok = true
	} else {
		ok = false
	}
	return
}

// returns true if the result has a previous page
func (r *GithubResult) HasPrevPage() (ok bool) {
	if (r.prevPageUrl != "") {
		ok = true
	} else {
		ok = false
	}
	return
}

// returns true if the result has a first page link
func (r *GithubResult) HasFirstPage() (ok bool) {
	if (r.firstPageUrl != "") {
		ok = true
	} else {
		ok = false
	}
	return
}

// returns true if the result has a last page link
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

// request the next page and return a new GithubResult or error (ErrNoNextPageUrl)
func (r *GithubResult) GetNextPage() (res *GithubResult, err error) {
	if !r.HasNextPage() {
		return nil, ErrNoNextPageUrl
	}

	return r.getPage(r.nextPageUrl)
}

// request the previous page and return a new GithubResult or error (ErrNoPrevPageUrl)
func (r *GithubResult) GetPrevPage() (res *GithubResult, err error) {
	if !r.HasPrevPage() {
		return nil, ErrNoPrevPageUrl
	}

	return r.getPage(r.prevPageUrl)
}

// request the first page and return a new GithubResult or error (ErrNoFirstPageUrl)
func (r *GithubResult) GetFirstPage() (res *GithubResult, err error) {
	if !r.HasFirstPage() {
		return nil, ErrNoFirstPageUrl
	}

	return r.getPage(r.firstPageUrl)
}

// request the last page and return a new GithubResult or error (ErrNoLastPageUrl)
func (r *GithubResult) GetLastPage() (res *GithubResult, err error) {
	if !r.HasLastPage() {
		return nil, ErrNoLastPageUrl
	}

	return r.getPage(r.lastPageUrl)
}
