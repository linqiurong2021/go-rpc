package main

import (
	"fmt"
	"net/rpc"
)

func main()  {
	// 1 建立连接
	client, err := rpc.Dial("tcp",":8800")
	if err != nil {
		fmt.Print("Dial error",err)
		return
	}
	defer client.Close()
	// 2: 传入参数 与 传出参数
	inputParams := "Clone"
	var outParams string
	// 2 调用
	err = client.Call("Hello.SayHello",inputParams,&outParams)
	if err != nil {
		fmt.Print("Call error",err)
		return
	}
	fmt.Println(outParams)
}
