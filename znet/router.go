package znet

import (
	"fmt"
	"github.com/zarttic/zinx/ziface"
)

// BaseRouter 实现路由时需要继承的基类，需要的时候对基类重写
type BaseRouter struct {
}

func (b *BaseRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call BaseRouter PreHandle")
}

func (b *BaseRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call BaseRouter Handle")
}

func (b *BaseRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call BaseRouter PostHandle")
}
