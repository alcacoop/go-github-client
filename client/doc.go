// Copyright 2012 Alca Società Cooperativa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package "client" implements a generic github api client.

It features Basic and OAuth2 authentication methods and abstract
pagination on github api results.

A simple usage example of a bare github client:

        ghc, _ := ghclient.NewGithubClient("username", "password", 
		                                   ghclient.AUTH_USER_PASSWORD)
       	req, _ := ghc.NewAPIRequest("GET", "user", nil)
        res, _ := ghc.RunRequest(req, new(http.Client))

        fmt.Println("JSON: %v\n", res.Json())


An example of paginated results:

        req, _ := ghc.NewAPIRequest("GET", "gists/public", nil)
        res, _ := ghc.RunRequest(req, new(http.Client))
        fmt.Println("PAGE 1: %v\n", res.Json())
        res, _ := res.GetNextPage();
        fmt.Println("PAGE 2: %v\n", res.Json())


*/
package client
