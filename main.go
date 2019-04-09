package main

import (
    "log"
    "net/http"

    "chat-demo/chat"
)

func main() {
    log.SetFlags(log.Lshortfile)

    // websocket server
    server := chat.NewServer("/entry")
    go server.Listen()

    log.Fatal(http.ListenAndServe(":3000", nil))
}
