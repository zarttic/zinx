package utils

import (
	"encoding/json"
	"github.com/zarttic/zinx/ziface"
	"os"
)

// GlobalObj 全局参数
type GlobalObj struct {
	// 服务器
	TcpServer ziface.IServer `json:"tcpServer"` //全局serve对象
	Host      string         `json:"host"`      //服务器监听的ip
	TcpPort   int            `json:"tcpPort"`   //服务器监听的端口
	Name      string         `json:"name"`      //当前服务器名称
	//Zinx
	Version        string `json:"version"`        //当前Zinx版本号
	MaxConnection  uint32 `json:"maxConn"`        //当前最大的连接数
	MaxPackageSize uint32 `json:"maxPackageSize"` //当前数据包的最大值
}

var GlobalObject *GlobalObj

// init 初始化 默认配置
func init() {
	GlobalObject = &GlobalObj{
		Name:           "ZinxServerApp",
		Version:        "V0.4",
		TcpPort:        8999,
		Host:           "0.0.0.0",
		MaxPackageSize: 4096,
		MaxConnection:  1000,
	}
	GlobalObject.Reload()
}
func (g *GlobalObj) Reload() {
	file, err := os.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(file, g)
	if err != nil {
		panic(err)
	}
}
