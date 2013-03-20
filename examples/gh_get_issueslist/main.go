// Copyright 2012 Alca Società Cooperativa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	//	"encoding/json"
	"flag"
	"fmt"
	ghclient "github.com/fluffle/go-github-client/client"
	ghissues "github.com/fluffle/go-github-client/issues"
)

func main() {
	help := flag.Bool("help", false, "Show usage")
	username := flag.String("u", "", "Specify Github user")
	password := flag.String("p", "", "Specify Github password")

	flag.Usage = func() {
		fmt.Printf("Usage:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *help == true || *username == "" || *password == "" {
		flag.Usage()
		return
	}

	ghc, err := ghclient.NewGithubClient(*username, *password, ghclient.AUTH_USER_PASSWORD)

	issuesc := ghissues.NewIssues(ghc)

	res, err := issuesc.GetRepoIssuesList("remogatto", "livegist",
		&ghissues.RepoListOptions{})

	fmt.Printf("RESPONSE: %v\nERROR: %v\n", res.RawHttpResponse, err)

	/*var body []byte = make([]byte, 10000)

		_, err = res.RawHttpResponse.Body.Read(body)*/

	//data, err := ioutil.ReadAll(res.RawHttpResponse.Body)

	jr, err := res.Json()

	/*var jr interface{}

		err := json.Unmarshal(([]byte)("[{\"prova\": 5}]"), &jr)*/

	fmt.Printf("JSON: %v\nERROR: %v\n", jr, err)
}
