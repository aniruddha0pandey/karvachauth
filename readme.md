# Karvachauth
A command line utility to add collaborators to GitHub repositories.

## Prequisits
Download [Go](https://golang.org/dl/). Setup the [GOPATH](https://github.com/golang/go/wiki/SettingGOPATH) variable to `export PATH=$PATH:$(go env GOPATH)/bin`.

## [Setting Project Layout](https://github.com/golang-standards/project-layout)
```
$ mkdir karvachaut 
$ git init && hub create
$ touch karvachaut.go readme.md
$ go mod init github.com/aniruddha0pandey/karvachauth 
$ go mod tidy && go build && go test
$ git add . && git commit -am "initial commit" && git push
$ go install github.com/aniruddha0pandey/karvachauth 
```

## [Setting Github Tokens](https://github.com/settings/tokens/new)
Select all scopes for `repo`, `admin:org`, `notifications` and `user` with note `Access Token - karvachauth`.
```
$ echo "{paste token from Github here excluding curly braces}" >> GITHUB_TOKEN
$ echo "GITHUB_TOKEN" >> .gitignore
```

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
