package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var mapApiDomainRoot = "https://api.openstreetmap.org/api/0.6/"

// 車関連 - 地図編

// 地図内の要素一つを表している
type MapElement struct {
	Type      string  `json:"type"`
	Id        int64   `json:"id"`
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
	TimeStamp string  `json:"timestamp"`
	Version   int16   `json:"version"`
	ChangeSet int64   `json:"changeset"`
	User      string  `json:"user"`
	Uid       int64   `json:"uid"`
}

// ダウンロード時のデータ取得
type FullMapResult struct {
	VersionInfo string       `json:"version"`
	Generator   string       `json:"generator"`
	CopyRight   string       `json:"copyright"`
	Attribution string       `json:"attribution"`
	License     string       `json:"license"`
	Elements    []MapElement `json:"elements"`
}

func main() {
	println("map data analize start")

	// --header 'Content-Type: application/json' \

	var url = mapApiDomainRoot + "relation/10000000/full.json"
	println(url)
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	resp, _ := client.Do(req)
	body, _ := io.ReadAll(resp.Body)

	var mapResult FullMapResult

	println(string(body))

	err := json.Unmarshal(body, &mapResult)
	if err != nil {
		print(err.Error())
		return
	}

	println("map analize complete")

	// 地図解析結果を出力
	println(mapResult.VersionInfo)
	println(mapResult.Generator)
	println(mapResult.License)
	println(mapResult.Attribution)
	println(mapResult.CopyRight)

	println("取得した要素一覧を取得")
	for _, element := range mapResult.Elements {
		println(element.Id)
		println(fmt.Sprintf("%f", element.Lat))
		println(fmt.Sprintf("%f", element.Lon))
	}

	// データ保存処理を開始(クライアントアプリ向けに整形をする)

}
