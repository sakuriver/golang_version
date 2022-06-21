package main

import "fmt"

// 統計の用語
func main() {

	score := 70
	averageValue := 10

	println(fmt.Sprintf("平均 %d 正規分布下限 %d 正規分布上限 %d", score, score-averageValue, score+averageValue))

}
