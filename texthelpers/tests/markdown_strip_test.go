package main

import (
	"github.com/fvbock/blackfriday"
	"fmt"
)

func main() {
	var md string
	md = `
## 1. Install go

They are two ways to install go.

1. Easist way:

  """
 brew install go
"""

2. Download from the [official page](http://golang.org/doc/install)

## 2. Setup GOROOT and GOPATH

1. Add the following line to your ~/.zshrc or ~/.bashrc

  """
 export GOROOT=/usr/local/Cellar/go/1.0.2/
 export GOPATH=$HOME/go
"""

2. NOTE: If you install go from downloaded package, set the _GOROOT_ to __/usr/local/go/__; for _GOPATH_, you can choose any path you like.

3. Setup code architecture.

  """
 cd $GOPATH; mkdir src bin pkg
"""

## 3. Install dependences

### Install brew

* mongodb

  """
  brew install mongodb
"""

* redis

  """
  brew install redis
"""

* bzr

  """
  brew install bzr
"""

* hg

  """
  brew install hg
"""

## 1~3. Install go and qortex dependencies for Linux (debian, ubuntu)

  """
  apt-get install golang mongodb bzr mercurial redis-server
"""

## 4. Checkout qortex

1. Create project diretory

  """
 cd $GOPATH; mkdir -p src/github.com/theplant/;
"""

2. Pull source code

  """
 cd src/github.com/theplant/
 git clone git@github.com:theplant/qortex.git
"""

## 5. Run

1. Get related go packages

  """
 cd $GOPATH/src/github.com/theplant/qortex/;
 go get .
"""

2. Start to run

  """
 go run main.go
"""

3. If it shows no error, it should be ok

  """
 ➜  qortex git:(master) ✗ go run main.go
2012/09/04 16:54:49 Starting server on :5000
"""
## 6. Load sample data

 Send @Aaron Liang or @Yeer Kunst your public key, and after they put your key on the server, then run these commands:

  """
  ./restore_without_files_from_dev.sh
  ./reset_local_possword.sh
  """

Then use one of these user to login: aaron@theplant.jp, password: nopassword

## 7. Migration

Remember to write down the migration history.
see __migration_history.md__



Running those creates your search index in redis. (maybe also other things? @Aaron Liang pls edit accordingly...)

  """
  $ cd $project_root/migrations/reindex
$ go run main.go -orgid=5018d345558fbe46c4000001
"""

The orgid passing in is the organization id, here is ThePlant.

Be aware that this will "FLUSHDB" your db 0 in redis... Make sure you don't have any important stuff there ;)

## 8. Deploy

1. Deploy on Development Server

  """
 ./deploy_by_build_server.sh
"""

2. Deploy on Production Server

## 9. Other tips

1. Setup alias to access the project directory
2. Setup alias go start the program
`

	sr := blackfriday.StrippedRenderer(0)
	extensions := 0
	fmt.Println(string( blackfriday.Markdown([]byte(md), sr, extensions) ))
}