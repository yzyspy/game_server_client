package main

import (
	"fmt"
	"io"
	"net"
	"time"
	"zinx-lib/zpack"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8877")
	if err != nil {
		panic(err)
	}
	data := []byte("hello-zinx")
	dp := zpack.DataPack{}

	for {
		message := zpack.NewMsgPackage(100, data)
		pack, err := dp.Pack(message)
		if _, err := conn.Write(pack); err != nil {
			fmt.Println("write err:", err)
			continue
		}

		binaryHead := make([]byte, dp.GetHeadLen())
		//前面8个字节是消息头,(int32 DateLen + int32 MsgID , 一共8个字节)
		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			fmt.Println("read head err:", err)
			continue
		}

		msgHead, err := dp.Unpack(binaryHead)
		if err != nil {
			fmt.Println("unpack head err2:", err)
			continue
		}

		fmt.Println("recv msg , msgId = %s, dataLen = %s", msgHead.GetMsgID(), msgHead.GetDataLen())

		msg := msgHead.(*zpack.Message)

		//再根据DateLen进行第二次读取, 将data取出来
		msg.Data = make([]byte, msgHead.GetDataLen())
		if _, err := io.ReadFull(conn, msg.Data); err != nil {
			fmt.Println("read data err:", err)
			continue
		}

		fmt.Println("recv data = ", string(msg.Data))

		time.Sleep(1000 * time.Millisecond)
	}
}
