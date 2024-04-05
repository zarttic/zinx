package znet

import (
	"fmt"
	"github.com/zarttic/zinx/ziface"
	"net"
)

// Server 接口实现，定义一个server的服务器模块
type Server struct {
	Name      string // 服务器名称
	IPVersion string // tcp4/tcp6
	IP        string // ip地址
	Port      int    // 端口
}

// Start 启动服务器
func (s *Server) Start() {
	fmt.Printf("[Zinx] Server Start IP %s, Port: %d \n", s.IP, s.Port)
	go func() {
		// 获取tcp的地址
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error: ", err)
			return
		}
		// 监听服务器地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen error: ", err)
			return
		}
		//阻塞的等待客户端链接，处理业务
		fmt.Println("start Zinx server success, ", s.Name, " success, listening...")
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept error: ", err)
				continue
			}
			// 已经建立链接，可以开始业务了
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("read error: ", err)
						continue
					}
					fmt.Printf("recv client data: %s, cnt: %d\n", buf, cnt)
					// 回显
					_, err = conn.Write(buf[:cnt])
					if err != nil {
						fmt.Println("write back error: ", err)
						continue
					}
				}
			}()

		}
	}()
}

// Stop 停止
func (s *Server) Stop() {
	// 释放资源
}

// Serve 运行服务器
func (s *Server) Serve() {
	//启动服务
	s.Start()
	//做一些启动后操作
	// 阻塞
	select {}
}

// NewServer 初始化server
func NewServer(name string) ziface.IServer {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
}
