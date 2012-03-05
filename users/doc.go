// Copyright 2012 Alca Società Cooperativa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package "users" implements a users github api client.

You need to create a github client and use it to create a new
users github client.

As an examples:

  ghc, err := ghclient.NewGithubClient(*username, *password, 
                                       ghclient.AUTH_USER_PASSWORD)
  ghusersc := ghusers.NewUsers(ghc)
  res, err := ghusersc.GetUserInfo(*userinfo)
  jr, err := res.Json()

  fmt.Printf("JSON: %v\nHTTP REPLY: %v\nERROR: %v\n", jr, res.RawHttpResponse, err)

*/
package users