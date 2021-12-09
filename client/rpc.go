package main

import (
	"os"
	"strconv"
)

type CallArg struct {
	Arg int
}

type CallReply struct {
	Ret int
	ID  int
}

func Sock() string {
	s := "/var/tmp/test-"
	s += strconv.Itoa(os.Getuid())
	return s
}
