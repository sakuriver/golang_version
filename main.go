package main

import (
	"fmt"
	"time"

	"./standard"
)

// ダンジョン突入前イメージ
func main() {
	partylists := standard.GetParties()

	for _, v := range partylists {
		fmt.Print(fmt.Sprintf("character id %d character name %s \r\n", v.Id, v.Name))
		fmt.Print(fmt.Sprintf("転生情報 番号 %d 名前 %s \r\n", v.Id, v.Name))
		time.Sleep(2)

	}
}
