package main

import "zinx-server/znet"

func main() {
	s := znet.NewServer("zinx server")
	s.Serve()
}
