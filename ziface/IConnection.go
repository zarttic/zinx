package ziface

import "net"

type IConnection interface {
	Start()                               // 启动
	Stop()                                // 结束
	GetTCPConnection() *net.TCPConn       // 获取连接
	GetConnID() uint32                    // 获取连接ID
	RemoteAddr() net.Addr                 // 获取远程地址
	Send(msgId uint32, data []byte) error // 发送信息

}

// HandleFunc 处理连接业务
type HandleFunc func(*net.TCPConn, []byte, int) error
