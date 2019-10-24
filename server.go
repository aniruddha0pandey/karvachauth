package main

import (
    "fmt"
    "log"
    "net/http"
)

var baseURL := "https://api.github.com"
var code string

var (
	ClientID := "8d95898e238907f7f66f"
	ClientSecret := "cb262234f230270d8235371749a16c9b582f54b5"
	CallbackURL := "http://localhost:6969/oauth/callback"
)

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
