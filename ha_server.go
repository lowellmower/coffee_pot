package main

import (
  "io"
  "log"
  "net/http"
  "fmt"
  "os/exec"
)

type handler func(w http.ResponseWriter, r *http.Request)

// handler filters
func GetOnly(h handler) handler {

    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "GET" {
            h(w, r)
            return
        }
        http.Error(w, "get only", http.StatusMethodNotAllowed)
    }
}

func PostOnly(h handler) handler {

    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            h(w, r)
            return
        }
        http.Error(w, "post only", http.StatusMethodNotAllowed)
    }
}

// Auth
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

// Routes
func HandleState(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Server is up.\n\n/on => device start\n\n/off => device stop\n\n")
}

func HandleOn(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Running 'On' script...\n")

    cmd := exec.Command("/bin/bash", "pi_on.sh")

    err := cmd.Start()
    if err != nil {
        log.Fatal(err)
    }
}

func HandleOff(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Running 'Off' script...\n")

    cmd := exec.Command("/bin/bash", "pi_off.sh")

    err := cmd.Start()
    if err != nil {
        log.Fatal(err)
    }
}


func main() {
    // get current status of device
    http.HandleFunc("/state", GetOnly(HandleState))

    // turn device on
    http.HandleFunc("/on", GetOnly(HandleOn))

    // turn device off
    http.HandleFunc("/off", GetOnly(HandleOff))

    log.Fatal(http.ListenAndServe(":8001", nil))
    fmt.Println("Listening on port: 8001\n")
}
