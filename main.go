package main

import (
	"github.com/MarkTBSS/074_Middleware/config"
	"github.com/MarkTBSS/074_Middleware/server"
)

func main() {
	conf := config.ConfigGetting()
	server := server.NewEchoServer(conf)
	server.Start()
}
