// Copyright 2012 Alca Società Cooperativa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package "gists" implements a gists github api client.

You need to create a github client and use it to create a new
gists github client.

As an examples:

  ghc, err := ghclient.NewGithubClient(*username, *password, 
                                       ghclient.AUTH_USER_PASSWORD)
  ghgistsc := ghgists.NewGists(ghc)
  res, err := ghgistsc.GetPublicGistsList()
  jr, err := res.Json()

  fmt.Printf("JSON: %v\nHTTP REPLY: %v\nERROR: %v\n", jr, res.RawHttpResponse, err)

  print("\n\nLOADING NEXT PAGE...\n\n")
  res, err = res.GetNextPage()

  jr, err = res.Json()

  fmt.Printf("JSON: %v\nERROR: %v\n", jr, err)
*/
package gists
