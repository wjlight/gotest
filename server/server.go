package main

import (
	"fmt"
	"io"
	"net"
)

const RECV_BUF_LEN = 1024

func main() {
	conOper()
}

func conOper() {
	lister, err := net.Listen("tcp", "127.0.0.1:6666")
	if err != nil {
		panic("error listening:" + err.Error())
	}
	fmt.Println("staring the server")

	for {
		conn, err := lister.Accept() //接受连接
		if err != nil {
			panic("Error accept:" + err.Error())
		}
		fmt.Println("Accepted the Connection:", conn.RemoteAddr())
		go EchoServer(conn)
	}
}

func EchoServer(conn net.Conn) {
	buf := make([]byte, RECV_BUF_LEN)

	defer conn.Close()

	for {
		n, err := conn.Read(buf)
		switch err {
		case nil:
			conn.Write(buf[0:n])
		case io.EOF:
			fmt.Println("End of data", err)
			return
		default:
			fmt.Println("Error:reading data")
			return
		}
	}
}
