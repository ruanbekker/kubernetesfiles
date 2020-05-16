package main

import (
    "fmt"
    "os"
    "net/http"
)

func hostnameHandler(w http.ResponseWriter, r *http.Request) {
    myhostname, _ := os.Hostname()
    mycontent     := "home"
    fmt.Fprintf(w, "content: %s, served from: %s \n", mycontent, myhostname)
}

func main() {
    const port string = "8000"
    fmt.Println("Server listening on port", port)
    http.HandleFunc("/", hostnameHandler)
    http.ListenAndServe(":" + port, nil)
}
