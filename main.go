package main

import (
    "context"
    "fmt"
    "os"
    "log"
    // "net/http"

    "golang.org/x/oauth2"
    // githuboauth  "golang.org/x/oauth2/github"
    "github.com/google/go-github/v28/github"

    "github.com/joho/godotenv"
)

func main() {
    // Load Environment Variables
    if godotenv.Load() != nil { 
        log.Fatal("Error loading .env file") 
    }

    // oauthStateString := os.Getenv("OAUTH_STATE")
    // oauthConf := &oauth2.Config{
    //     ClientID:     os.Getenv("CLIENT_ID"),
    //     ClientSecret: os.Getenv("CLIENT_SECRET"),
    //     Scopes:       []string{"admin:org", "repo", "user", "notifications"},
    //     Endpoint:     githuboauth.Endpoint,
    // }

    ctx := context.Background()
    ts := oauth2.StaticTokenSource( &oauth2.Token{ AccessToken: os.Getenv("ACCESS_TOKEN") } )
    tc := oauth2.NewClient(ctx, ts)

    client := github.NewClient(tc)

    repos, _, err := client.Repositories.List(ctx, "", nil)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(repos)
    
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



// tmp access_token=4b78a2da2ac65ad8e8b3407befe8e61e551420d5

