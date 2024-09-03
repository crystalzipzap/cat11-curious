package server

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/crystalzipzap/cat11-curious/decoder/internal/message"
)

// StartMockServer generates and sends Category 11 messages
func StartMockServer() {
    listener, err := net.Listen("tcp", ":8080")
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
        go sendCategory11Messages(conn)
    }
}

func sendCategory11Messages(conn net.Conn) {
    defer conn.Close()
    encoder := json.NewEncoder(conn)

    for {
        message := message.Category11Message{
            MessageType: "Category11",
            TimeOfDay:   time.Now().Format(time.RFC3339),
            SensorID:    "Sensor1",
        }

        err := encoder.Encode(message)
        if err != nil {
            fmt.Println("Error sending message:", err)
            return
        }
        time.Sleep(1 * time.Second)
    }
}