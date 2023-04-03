package main

import (
	"fmt"
	"net"
)

func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

func main() {
	free_port, err := GetFreePort()
	if free_port != 0 {
		fmt.Println("Free port is: ", free_port)
	} else {
		fmt.Println("Free port is not available. Error is: ", err)
	}
}
