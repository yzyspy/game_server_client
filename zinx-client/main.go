package main

import (
	"fmt"
	"net"
	"zinx-lib/zpack"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8877")
	if err != nil {
		panic(err)
	}
	data := []byte("hello-zinx")

	dataPack := zpack.DataPack{}
	message := zpack.NewMsgPackage(1223, data)
	pack, err := dataPack.Pack(message)
	conn.Write(pack)

	//conn.Write(data)

	buf := make([]byte, 1024)
	if _, err := conn.Read(buf); err != nil {
		panic(err)
	}
	fmt.Println(string(buf))
}
