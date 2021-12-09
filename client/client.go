package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	callarg := CallArg{}
	callarg.Arg = 1
	reply := CallReply{}
	reply.Ret = 1
	reply.ID = 1
	sockname := Sock()
	c, err := rpc.DialHTTP("unix", sockname)
	if err != nil {
		log.Fatal("DialHttp error:", err)
	}
	defer c.Close()
	err = c.Call("Server.Request", &callarg, &reply)
	if err != nil {
		log.Fatal("Call error:", err)
	}
	fmt.Println(reply.Ret)
	fmt.Println(reply.ID)
	reply.ID = 1
	callarg.Arg = 0
	err = c.Call("Server.Request", &callarg, &reply)
	if err != nil {
		log.Fatal("Call error:", err)
	}
	fmt.Println(reply.Ret)
	fmt.Println(reply.ID)
}
