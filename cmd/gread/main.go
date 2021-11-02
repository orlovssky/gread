package main

import (
	"github.com/orlovssky/gread/internal/server"
)

func main() {
	server := server.NewServer()
	server.StartServer()
}
