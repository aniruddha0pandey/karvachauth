# Karvachauth
A command line utility to add collaborators to GitHub repositories.

## Prequisits
- Install [Hub](https://github.com/github/hub#installation).
- Install [Go](https://golang.org/dl/).
- Setup the [GOPATH](https://github.com/golang/go/wiki/SettingGOPATH) variable to `export PATH=$PATH:$(go env GOPATH)/bin`.

## [Setting Project Layout](https://github.com/golang-standards/project-layout)
```
$ go get github.com/spf13/cobra/cobra
$ mkdir karvachauth && cd ./karvachauth
$ cobra init --pkg-name github.com/aniruddha0pandey/karvachauth
$ declare -a commands=("login" "repos" "collab" "add" "remove")
$ for i in "${commands[@]}"; do cobra add $i; done
$ touch server.go README.md
$ go mod init github.com/aniruddha0pandey/karvachauth
$ go build ./bin
$ go install github.com/aniruddha0pandey/karvachauth
$ echo "bin/" >> .gitignore
$ hub init && hub create
$ hub commit -am "initial commit" && hub push
```

## [Setting OAuth Client](https://github.com/settings/developers)
|---|---|
| Scopes | `repo` `admin:org` `notifications` `user` |
| Homepage URL | http://localhost:6969/home |
| Authorization callback URL | http://localhost:6969/oauth/callback |

## Usage
```
$ karvachauth --help
NAME:
	karvachauth - A command line utility to add collaborators to GitHub repositories

USAGE:
	karvachauth [global options] command [command options] [arguments...]

VERSION:
	0.0.1

COMMANDS:
	help			Shows a list of commands or help for one command
	login			Enter username password to OAuth token
	repos			list all current repositories
	collab			list all current collabrators
	add				add collabrators to repositories
	remove			remove collabrators from repositories

GLOBAL OPTIONS:
	--help, -h		show help
	--version, -v	print the version
```
