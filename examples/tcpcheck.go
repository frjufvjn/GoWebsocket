package main

import (
	"fmt"
	"net"
)

func main() {
	isAlive := tcpAliveCheck()
	fmt.Println("isAlive : ", isAlive)
}

func tcpAliveCheck() bool {
	_, err := net.Dial("tcp", "localhost:18080")
	if nil != err {
		fmt.Println("failed to connect to server")
		return false
	} else {
		return true
	}
}
