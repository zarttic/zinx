package ziface

// IDataPack 处理tcp粘包问题
type IDataPack interface {
	GetHeadLen() uint32                         //获取包头长度
	Pack(msg IMessage) ([]byte, error)          //封包
	Unpack(binaryData []byte) (IMessage, error) //拆包
}
