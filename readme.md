# Karvachauth
A command line utility to add collabrators to github repositories.

## Prequisits
Download [Go](https://golang.org/dl/). Setup the [GOPATH](https://github.com/golang/go/wiki/SettingGOPATH) variable to `$(go env GOPATH)/bin`.

```
$ mkdir karvachaut 
$ git init && hub create
$ touch karvachaut.go readme.md
$ go mod init github.com/aniruddha0pandey/karvachauth 
$ go build && go mod tidy
$ git add . && git commit -am "initial commit" && git push
$ karvachauth help
usage: karvachauth [command] <args>

login		Enter username password to generate token
repos		list all current repositories
collab		list all current collabrators
add			add collabrators to repositories

```