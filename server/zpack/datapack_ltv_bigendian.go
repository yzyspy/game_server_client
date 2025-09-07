package zpack

import (
	"zinx-server/ziface"
)

type DataPack struct{}

func (d *DataPack) GetHeadLen() uint32 {
	//TODO implement me
	panic("implement me")
}

func (d *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DataPack) Unpack(bytes []byte) (ziface.IMessage, error) {
	//TODO implement me
	panic("implement me")
}
