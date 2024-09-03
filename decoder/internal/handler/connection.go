package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net"

	"github.com/crystalzipzap/cat11-curious/decoder/internal/message"
)

// HandleConnection handles the client connection and decodes messages
func HandleConnection(conn net.Conn) {
    defer conn.Close()
    decoder := json.NewDecoder(conn)

    for {
        var msg message.Category11Message
        err := decoder.Decode(&msg)
        if err != nil {
            if err == io.EOF {
                fmt.Println("Connection closed by the client")
            } else {
                fmt.Println("Error decoding message:", err)
            }
            return
        }

        fmt.Printf("Received message: %+v\n", msg)

        parsedMessage, err := message.ParseCategory11Message(msg.MessageType)
        if err != nil {
            fmt.Println("Error parsing message:", err)
            continue
        }

        fmt.Printf("Parsed message: %+v\n", parsedMessage)
    }
}