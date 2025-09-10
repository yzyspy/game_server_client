package main

import "zinx-lib/znet"

func main() {
	s := znet.NewServer("zinx server")

	router := znet.EchoRouter{}
	s.AddRouter(100, &router)

	s.Serve()
}
