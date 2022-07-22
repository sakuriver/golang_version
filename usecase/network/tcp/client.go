package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "docs.studygolang.com:80")
	if err != nil {
		println(err.Error())
		return

	}
	println(conn.RemoteAddr().String())

	fmt.Fprintf(conn, "GET / Http/1.0")
}
