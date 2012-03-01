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
