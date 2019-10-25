package main

import (
    "fmt"
    "log"
    "net/http"
)

var AccessCode string
var AccessToken string

// var (
// 	ClientID := "8d95898e238907f7f66f"
// 	ClientSecret := "cb262234f230270d8235371749a16c9b582f54b5"

//     baseURL := "https://github.com"
// 	CallbackURL := "http://localhost:6969/oauth/callback"
// )

func handleHome(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s/\n", r.URL.Path[1:])
    fmt.Fprintf(w, "Karvachauth")
}

func handle404(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s/\n", r.URL.Path[1:])
    fmt.Fprintf(w, "Karvachauth")
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s/\n", r.URL.Path[1:])
    fmt.Fprintf(w, "Karvachauth")
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s/\n", r.URL.Path[1:])
    fmt.Fprintf(w, "Karvachauth")
}

func handleSuccess(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%s/\n", r.URL.Path[1:])
    fmt.Fprintf(w, "Karvachauth")
}

func main() {
    http.HandleFunc("/home", handleHome)
    http.HandleFunc("/404", handle404)
    http.HandleFunc("/login", handleLogin)
    http.HandleFunc("/oauth/callback", handleCallback)
    http.HandleFunc("/oauth/success", handleSuccess)
    log.Fatal(http.ListenAndServe(":6969", nil))
}


// GET baseURL/login/oauth/authorize?client_id=ClientID&redirect_uri=CallbackURL
// GET CallbackURL?code=f867948d80c5ceee1b6c
// POST baseURL/login/oauth/access_token?client_id=ClientID&client_secret=ClientSecret?code=AccessCode
// GET access_token=0d490cd3e30f2be1e8c28529c9042f040059ff6a&scope=&token_type=bearer


