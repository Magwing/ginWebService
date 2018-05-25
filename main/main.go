package main

import (
	_ "github.com/Wan-Mi/ginWebService/router"
	"github.com/Wan-Mi/ginWebService/server"
)

func main() {
	httpServer := server.NewServer(":8989")
	server.Run(httpServer)
}
