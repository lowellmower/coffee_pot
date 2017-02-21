package haserver

import (
    "encoding/base64"
    "net/http"
    "strings"
)

type handler func(w http.ResponseWriter, r *http.Request)

func BasicAuth(pass handler) handler {

    return func(w http.ResponseWriter, r *http.Request) {
        username, password, _ := r.BasicAuth()

        if username != "username" || password != "password" {
            http.Error(w, "authorization failed", http.StatusUnauthorized)
            return
        }

        pass(w, r)
    }
}
