package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type Hello struct {
}

// SayHello 必须是可导出的(外部可访问)
// 第一个参数 基本类型或map slice
// 第二个参数必须是指针
// 只有一个返回参数 error
func (hello *Hello) SayHello(name string, resp *string) error {
	*resp = name + "您好"
	return nil
}

func main() {
	// 1 注册服务
	rpc.RegisterName("Hello", &Hello{})
	// 2 监听
	listener, err := net.Listen("tcp", ":8800")
	if err != nil {
		fmt.Print("Listen err:", err)
		return
	}
	defer listener.Close()
	// 3 接收数据
	conn, err := listener.Accept()
	if err != nil {
		fmt.Print("Accept err:", err)
		return
	}
	defer conn.Close()
	// 4 监听
	rpc.ServeConn(conn)

}
