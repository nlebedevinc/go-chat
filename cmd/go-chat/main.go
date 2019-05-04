package main

import (
    "log"
    "net/http"

    "github.com/nlebedevinc/go-chat/pkg/chat"
)

func main() {
    log.SetFlags(log.Lshortfile)

    // websocket server
    server := chat.NewServer("/entry")
    go server.Listen()

    log.Fatal(http.ListenAndServe(":3000", nil))
}
