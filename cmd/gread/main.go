package main

import "github.com/orlovssky/gread/internal/server"

func main() {
	s := server.NewServer()
	s.StartServer()
}
