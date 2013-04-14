// Copyright 2012 Alca Società Cooperativa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	ghclient "github.com/iansmith/go-github-client/client"
	ghgists "github.com/iansmith/go-github-client/gists"
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

	ghc, _ := ghclient.NewGithubClient(*username, *password, ghclient.AUTH_USER_PASSWORD)

	ghc_gists := ghgists.NewGists(ghc)

	newGist := ghgists.NewGistDataCreate()
	newGist.Description = "test go create api"
	newGist.Public = false
	newGist.Files["prova1.js"] = ghgists.GistFileContent{Content: "var x=5;"}

	res, _ := ghc_gists.CreateGist(newGist)

	jr, _ := res.Json()
	gist_id := jr.GetString("id")

	data := ghgists.NewGistDataUpdate()
	data.Files["prova1.js"] = ghgists.GistFileContent{Filename: "renamed.js", Content: "var x=5;\nvar y=10;"}

	fmt.Printf("UPDATE DATA: %v\n", data)

	res, _ = ghc_gists.UpdateGist(gist_id, data)

	jr, _ = res.Json()

	fmt.Printf("JSON: %v\n", jr)
	fmt.Printf("RAW: %v\n", res.RawHttpResponse)
}
