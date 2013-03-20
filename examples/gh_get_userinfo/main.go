// Copyright 2012 Alca Società Cooperativa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	ghclient "github.com/fluffle/go-github-client/client"
	ghusers "github.com/fluffle/go-github-client/users"
)

func main() {
	help := flag.Bool("help", false, "Show usage")
	username := flag.String("u", "", "Specify Github user")
	password := flag.String("p", "", "Specify Github password")
	userinfo := flag.String("userinfo", "", "Specify another Github user")

	flag.Usage = func() {
		fmt.Printf("Usage:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	if *help == true || *username == "" || *password == "" || *userinfo == "" {
		flag.Usage()
		return
	}

	ghc, err := ghclient.NewGithubClient(*username, *password,
		ghclient.AUTH_USER_PASSWORD)

	ghusersc := ghusers.NewUsers(ghc)

	res, err := ghusersc.GetUserInfo(*userinfo)

	jr, err := res.Json()

	fmt.Printf("JSON: %v\nHTTP REPLY: %v\nERROR: %v\n", jr, res.RawHttpResponse, err)
}
