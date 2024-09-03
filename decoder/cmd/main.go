package main

import (
	"github.com/crystalzipzap/cat11-curious/decoder/internal/server"
)

func main() {
	go server.StartMockServer()

	server.StartTCPServer()
}