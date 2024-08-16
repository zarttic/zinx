package znet

// Message 请求的消息

type Message struct {
	Id   uint32 //消息id
	Len  uint32 //消息长度
	Data []byte //消息
}

func (m *Message) GetMsgId() uint32 {
	return m.Id
}

func (m *Message) GetMsgLen() uint32 {
	return m.Len
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetMsgId(u uint32) {
	m.Id = u
}

func (m *Message) SetMsgLen(u uint32) {
	m.Len = u
}

func (m *Message) SetData(data []byte) {
	m.Data = data
}
