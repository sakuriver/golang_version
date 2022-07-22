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
	// ダイアルした結果のipアドレス情報を取得
	println("リモート情報")
	println(conn.RemoteAddr().String())
	println("ローカルアドレス情報")
	println(conn.LocalAddr().String())

	fmt.Fprintf(conn, "GET / Http/1.0")
}
