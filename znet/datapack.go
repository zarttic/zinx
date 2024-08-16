package znet

import "github.com/zarttic/zinx/ziface"

type DataPack struct {
}

func NewDataPack() *DataPack {
	return &DataPack{}
}
func (d *DataPack) GetHeadLen() uint32 {
	//data_len(4) + id(4)
	return 8
}

func (d *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {
	panic("implement me")
}

func (d *DataPack) Unpack(binaryData []byte) (ziface.IMessage, error) {
	//TODO implement me
	panic("implement me")
}
