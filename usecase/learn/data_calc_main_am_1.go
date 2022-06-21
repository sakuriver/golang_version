package main

import "fmt"

// 機械向けのデータに変換
func main() {

	j := 50
	initialValue := j
	NISHIN := []int{}
	for i := 0; i < 8; i++ {
		NISHIN = append(NISHIN, j%2)
		j = j / 2

	}
	strResult := ""
	for i := 0; i < 8; i++ {
		strResult += fmt.Sprintf("%d", NISHIN[i])
		println(fmt.Sprintf("%d番目 %d", i, NISHIN[i]))
	}
	print(fmt.Sprintf("人間語 %d  機械語 %s", initialValue, strResult))

}
