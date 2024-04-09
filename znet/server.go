package znet

import (
	"fmt"
	"github.com/zarttic/zinx/utils"
	"github.com/zarttic/zinx/ziface"
	"net"
)

// Server 接口实现，定义一个server的服务器模块
type Server struct {
	Name      string         // 服务器名称
	IPVersion string         // tcp4/tcp6
	IP        string         // ip地址
	Port      int            // 端口
	Router    ziface.IRouter // 路由
}

func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
}

// CallBackDefault  回显
func CallBackDefault(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("[Conn Handle] CallBackDefault ... ")
	_, err := conn.Write(data[:cnt])
	if err != nil {
		fmt.Println("Write error: ", err)
		return err
	}
	fmt.Println("recv from client: ", string(data[:cnt]))
	return nil
}

// Start 启动服务器
func (s *Server) Start() {
	fmt.Println("[Zinx] Server Name: ", utils.GlobalObject.Name, ", Server Version: ", utils.GlobalObject.Version)
	fmt.Println("MaxConnection: ", utils.GlobalObject.MaxConnection, " MaxPackageSize: ", utils.GlobalObject.MaxPackageSize)
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
		var cid uint32 = 0

		//阻塞的等待客户端链接，处理业务
		fmt.Println("start Zinx server success, ", s.Name, " success, listening...")
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept error: ", err)
				continue
			}
			dealConn := NewConnection(conn, cid, s.Router)
			cid++
			//启动业务处理
			go dealConn.Start()
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
func NewServer() ziface.IServer {
	return &Server{
		Name:      utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.TcpPort,
		Router:    &BaseRouter{},
	}
}
