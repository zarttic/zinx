package znet

import (
	"fmt"
	"github.com/zarttic/zinx/utils"
	"github.com/zarttic/zinx/ziface"
	"net"
)

type Connection struct {
	Conn     *net.TCPConn   //套接字
	ConnID   uint32         //连接id
	isClosed bool           //是否关闭
	ExitChan chan bool      //告知当前连接已经退出
	Router   ziface.IRouter //当前链接的路由
}

// StartReader 读业务
func (c *Connection) StartReader() {
	fmt.Println("reader goroutine running [id] -> ", c.ConnID)
	defer fmt.Println("reader goroutine exit [id] -> ", c.ConnID)
	defer c.Stop()
	for {
		buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("read error -> ", err)
			c.ExitChan <- true
			continue
		}
		req := Request{
			conn: c,
			data: buf,
		}
		// 从路由绑定的处理函数
		go func(req ziface.IRequest) {
			c.Router.PreHandle(req)
			c.Router.Handle(req)
			c.Router.PostHandle(req)
		}(&req)

	}
}
func (c *Connection) Start() {
	fmt.Println("connection start [id] -> ", c.ConnID)
	//启动读协程
	go c.StartReader()
	for {
		select {
		case <-c.ExitChan:
			//得到退出消息，不再阻塞
			return
		}
	}
}

func (c *Connection) Stop() {
	if c.isClosed {
		return
	}
	c.isClosed = true
	defer c.Conn.Close()
	c.ExitChan <- true
	close(c.ExitChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
func (c *Connection) Send(msgId uint32, data []byte) error {
	//TODO implement me
	panic("implement me")
}

// NewConnection 创建连接
func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	return &Connection{
		Conn:     conn,
		ConnID:   connID,
		isClosed: false,
		Router:   router,
		ExitChan: make(chan bool, 1),
	}
}
