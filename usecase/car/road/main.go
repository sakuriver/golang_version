package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

var mapApiDomainRoot = "https://api.openstreetmap.org/api/0.6/"

// 車関連 - 道路編

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

// アプリ用の地図マスターデータと実行結果出力
func main() {
	println("map data analize start")

	startTime := time.Now()

	// --header 'Content-Type: application/json' \

	var url = mapApiDomainRoot + "relation/10000005/full.json"
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

	//元データが取得できたので、書き込むマスターデータを開く
	f, err := os.OpenFile("masterdata/map.json", os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModeAppend.Perm())
	if err != nil {
		println(err.Error())
		return
	}

	println("取得した要素一覧を取得")
	for _, element := range mapResult.Elements {
		println(element.Id)
		println(fmt.Sprintf("緯度 %f", element.Lat))
		println(fmt.Sprintf("経度 %f", element.Lon))
	}

	_, err = f.Write(body)
	if err != nil {
		println(err.Error())
		return
	}

	defer f.Close()

	// データ保存処理を開始(クライアントアプリ向けに整形をする)

	// アプリ内で使う配信データの生成時間を出力(国や地域拡大時の改善指標と運営レポートで利用)

	println("map analize time %v", time.Since(startTime))

	// 番号と実行時間によって「開発におけるリードタイム改善」用出力
	// ミリ秒と秒単位での改善レポート
	// ジョブズの言葉を借りると「１ユーザー当たり〇秒改善」及び「運営時のコスト〇秒改善により１日前の開発ストップを半日前に改善案」
	developDataFormat := fmt.Sprintf("%d,%v", 1, time.Since(startTime).Milliseconds())

	println(developDataFormat)
	// 別kpiバッチ予定 前回から見た場合の待機時間改善秒数
	records := [][]string{
		{"No", "ExecTime"},
		{"1", strconv.FormatInt(time.Since(startTime).Milliseconds(), 10)},
	}

	// 日付でのKPIローテーション
	csvfile, err := os.Create(fmt.Sprintf("./kpi/map_generate_time_%s.csv", time.Now().UTC().Format("2006-01-02")))

	w := csv.NewWriter(csvfile)

	w.WriteAll(records)

	// 経営層及び投資向けのレポーティングファイルとしてcsvファイル出力

	// 本ジョブ以外の開発コスト用レポーティング集計

	// １日の時間 - 間の時間(ここが0に近づくと、間の作業が0になる)

	w.Flush()
	csvfile.Close()

}
