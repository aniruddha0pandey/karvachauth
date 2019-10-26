package main

import (
    "bufio"
    "context"
    "fmt"
    "os"
    "strings"
    "syscall"

    "github.com/google/go-github/v28/github"
    "golang.org/x/crypto/ssh/terminal"
)

func main() {
    var username string
    fmt.Print("Enter GitHub username: ")
    fmt.Scanf("%s", &username)
}


// BaseURL = "https://github.com"
// ApiURL = "https://api.github.com"
// CallbackURL = "http://localhost:6969/oauth/callback"

// Revoke access BaseURL/settings/connections/applications/ClientID

// "Authorization: token AccessToken"

// ----------------------------------------------------------------------------

/* karvachauth login */
// GET BaseURL/login/oauth/authorize?client_id=ClientID&redirect_uri=CallbackURL
// GET CallbackURL?code=f867948d80c5ceee1b6c
// POST baseURL/login/oauth/access_token?client_id=ClientID&client_secret=ClientSecret?code=AccessCode
// GET access_token=0d490cd3e30f2be1e8c28529c9042f040059ff6a&scope=&token_type=bearer

// ----------------------------------------------------------------------------

/* karvachauth repos */
// GET ApiURL/user/repos

/* karvachauth --org repos ORGANIZATION-NAME */
// GET ApiURL/orgs/:org/repos

// ----------------------------------------------------------------------------

/* karvachauth collab REPOSITORY-NAME */
// GET ApiURL/repos/:owner/:repo/collaborators

// ----------------------------------------------------------------------------

/* karvachauth add REPOSITORY-NAME USER-NAME */
// PUT ApiURL/repos/:owner/:repo/collaborators/:username

// ----------------------------------------------------------------------------

/* karvachauth remove REPOSITORY-NAME USER-NAME */
// DELETE ApiURL/repos/:owner/:repo/collaborators/:username

// ----------------------------------------------------------------------------

/* karvachauth help */

// ----------------------------------------------------------------------------



// tmp access_token=ad1c3191ecc9be99f550b3499539915a0490fe5c

