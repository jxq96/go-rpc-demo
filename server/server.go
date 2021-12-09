package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

type Server struct{}

func (s *Server) Request(arg *CallArg, reply *CallReply) error {
	fmt.Println(arg.Arg)
	reply.Ret = 2
	reply.ID = 0
	return nil
}

func main() {
	server := Server{}
	rpc.Register(&server)
	rpc.HandleHTTP()
	sockname := Sock()
	os.Remove(sockname)
	l, err := net.Listen("unix", sockname)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	http.Serve(l, nil)
}
