// Copyright 2012 Alca Società Cooperativa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package "issues" implements a issues github api client.

You need to create a github client and use it to create a new
issues github client.

As an examples:

  ghc, err := ghclient.NewGithubClient(*username, *password, 
                                       ghclient.AUTH_USER_PASSWORD)
  ...
*/
package issues
