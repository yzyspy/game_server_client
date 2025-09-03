package zpack

type Message struct {
	DataLen uint32 // Length of the message
	ID      uint32 // ID of the message
	Data    []byte // Content of the message
	rawData []byte // Raw data of the message
}

func (m *Message) GetDataLen() uint32 {
	//TODO implement me
	panic("implement me")
}

func (m *Message) GetMsgID() uint32 {
	//TODO implement me
	panic("implement me")
}

func (m *Message) GetData() []byte {
	//TODO implement me
	panic("implement me")
}

func (m *Message) GetRawData() []byte {
	//TODO implement me
	panic("implement me")
}

func (m *Message) SetMsgID(u uint32) {
	//TODO implement me
	panic("implement me")
}

func (m *Message) SetData(bytes []byte) {
	//TODO implement me
	panic("implement me")
}

func (m *Message) SetDataLen(u uint32) {
	//TODO implement me
	panic("implement me")
}
