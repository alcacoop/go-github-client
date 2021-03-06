go-github-client: Github v3 API Client for Go
=============================================

DEPRECATED
----------

**go-github-client** is now superseded by [google/go-github](https://github.com/google/go-github)
package released by Google.

Description
-----------

go-github-client</tt> is a Go package that aims to smoothly
integrate Github Rest API in Go applications.

WARNING: This software is alpha quality so use it at your risks ;-)

We've extracted and refactored this package from
[cocode.io](http://cocode.io), our real-time web collaborative code
editor integrated with GitHub Gist. The package is currently under
active development because we're enhancing the integration between
cocode.io</tt> and GitHub.

To keep you updated about future development, please feel free to watch 
this project on GitHub and to follow us on:

* [@cocodeio (twitter)](https://twitter.com/#!/cocodeio)
* [+cocodeio (google+)](https://plus.google.com/110953439702828767840/posts)

Of course, we're looking forward for your feedback & patches :)
 
Quickstart Examples
-------------------

### go get & go install

Download go-gothub-client and install go-github-client/client:

```
$ go get github.com/alcacoop/go-github-client/client
```

This command clone all the repository in "$GOPATH/src/github.com/alcacoop/go-github-client"
and install "github.com/alcacoop/go-github-client/client" package in "$GOPATH/pkg/ARCH/github.com/alcacoop/go-github-client/client.a"

Now you can "go install" all the other subpackage (or just what you want):

```
$ go install github.com/alcacoop/go-github-client/users

$ go install github.com/alcacoop/go-github-client/gists

$ go install github.com/alcacoop/go-github-client/issues
```

### go-github-client command line example applications

You can build and install some example applications using "go install":

```
$ go install github.com/alcacoop/go-github-client/examples/gh_get_userinfo

$ $GOPATH/bin/gh_get_userinfo
...
```

### Create an authenticated GitHub API Client using Basic Authentication method

```
import (
  ghclient "github.com/alcacoop/go-github-client/client"
)

...

ghc, _ := ghclient.NewGitHubClient("testuser", "password", AUTH_USER_PASSWORD)

```  

### Get information about the authenticated user

```
import (
  ghusers "github.com/alcacoop/go-github-client/users"
)

usersc = ghusers.NewUsers(ghc)
res, err := usersc.GetUserInfo(*userinfo)
jr, err := res.Json()
name := jr.GetString("login")

```

### Paginate public gists 

```
import (
  ghgist "github.com/alcacoop/go-github-client/gists"
)

gistsc = ghgists.NewUsers(ghc)
res, err := gistsc.GetPublicGists()
jr, err := res.Json()

res2, err := res.NextPage()
jr2, err := res.Json()
```

More Info
-------------

### API Reference
* [Base Client](http://gopkgdoc.appspot.com/pkg/github.com/alcacoop/go-github-client/client)
* [Users](http://gopkgdoc.appspot.com/pkg/github.com/alcacoop/go-github-client/users)
* [Gists](http://gopkgdoc.appspot.com/pkg/github.com/alcacoop/go-github-client/gists)

### Examples
* [Usage Examples](https://github.com/alcacoop/go-github-client/tree/master/examples)



