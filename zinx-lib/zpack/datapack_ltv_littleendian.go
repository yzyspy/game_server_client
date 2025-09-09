package zpack

import "zinx-lib/ziface"

type DataPackLtv struct{}

func (d *DataPackLtv) GetHeadLen() uint32 {
	//TODO implement me
	panic("implement me")
}

func (d *DataPackLtv) Pack(msg ziface.IMessage) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DataPackLtv) Unpack(bytes []byte) (ziface.IMessage, error) {
	//TODO implement me
	panic("implement me")
}
