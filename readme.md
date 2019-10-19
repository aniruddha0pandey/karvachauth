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
Select all scopes for `repo`, `admin:org`, `notifications` and `user` with note `Access Token for karvachauth`.
```
$ echo "{paste token from Github here}" >> token
$ echo "token" >> .gitignore
```

## Usage
$ karvachauth --help
usage: karvachauth [command] <args>

login		Enter username password to generate token
repos		list all current repositories
collab		list all current collabrators
add			add collabrators to repositories
```

