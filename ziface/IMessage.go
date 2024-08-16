package ziface

// IMessage 请求消息抽象接口
type IMessage interface {
	GetMsgId() uint32  //获取消息Id
	GetMsgLen() uint32 //获取消息长度
	GetData() []byte   //获取消息内容
	SetMsgId(uint32)   //设置消息id
	SetMsgLen(uint32)  //设置消息长度
	SetData([]byte)    //设置消息内容
}
