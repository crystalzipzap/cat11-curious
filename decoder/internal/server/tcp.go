package server

import (
	"fmt"
	"net"

	"github.com/crystalzipzap/cat11-curious/decoder/internal/handler"
)

// StartTCPServer receives and decodes Category 11 messages
func StartTCPServer() {
    listener, err := net.Listen("tcp", ":9090")
    if err != nil {
        fmt.Println("Error creating server:", err)
        return
    }
    defer listener.Close()

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }
        go handler.HandleConnection(conn)
    }
}