package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

// 1ヵ月当たりでの時間 週5 8時間で経験者を数える
const MonthHour = 8 * 20

// 経験年数〇年　用の時間
const YearHour = MonthHour * 12

// 個人 - スキルシート集計結果編
// 年数ではなく粒度を把握するためを想定

// 基本情報データ
type PersonalBasicInformation struct {
	Id           int    `json:"id"`
	NickName     string `json:"nick_name"`
	FromName     string `json:"from_name"`
	RealLocation string `json:"real_location"`
}

// 業務用の住所情報
type LocationInformation struct {
	// 地方公共団体コード
	AreaCode int `json:"area_code"`
	// 郵便番号(旧)
	OldPostNum int `json:"old_post_num"`
	// 郵便番号
	PostNum string `json:"post_num"`
	// 都道府県名(カナ)
	PrefectureKana string `json:"prefecture_kana"`
	// 市区町村名(カナ)
	CityNameKana string `json:"city_name_kana"`
	// 町域名(カナ)
	TownNameKana string `json:"town_name_kana"`
	// 都道府県名
	PrefectureName string `json:"prefecture_name"`
	// 市区町村名
	CityName string `json:"city_name"`
	// 町域名
	TownName string `json:"town_name"`
}

// 業務のプロジェクト及び所属業務一覧
type TaskProject struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// 製品の特定ジャンルについて、詳しいかの確認用
type ProductGenre struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// 仕事で使っているメーカーコンピュータのカテゴリー一覧
type BenderCategory struct {
	Id         int    `json:"id"`
	BenderType int    `json:"bender_type"`
	Name       string `json:"name"`
}

// データ保存ソフトウェアの一覧
type DataSaveMiddleWare struct {
	Id       int    `json:"id"`
	Category int    `json:"category"`
	Name     string `json:"name"`
}

// プログラミング言語経験(ソフトウェアエンジニア関連固有テーブル)
type ProgrammingLanguage struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	TotalSum int    `json:"total_sum"`
}

// 業務詳細(企業固有にならない部分としての一覧)
type JobDescription struct {
	Id       int    `json:"id"`
	Category int    `json:"category"`
	Name     string `json:"name"`
}

// おかいものしたもの
type ShoppingHistory struct {
	// おみせ番号
	ShopId int `json:"shop_id"`
	// こうにゅうしたもの
}

// アプリ側で利用するコンテンツデータ
// パッケージデータとしてダウンロードできるように１個の構造体に設定
// 容姿ではなく、開発プロジェクトにジョインするための内容として集計
type JocContentsDatabase struct {
	// プロジェクト種類 大枠での集計で利用
	TaskProjects []TaskProject `json:"task_projects"`
	// 開発してきた製品のジャンル
	//ProductGenres []ProductGenre `json:"genres"`
	// 特定製品系強い系の確認
	//Benders []BenderCategory `json:"benders"`
	// ミドルウェアの利用経験年数(sdkとかデータベースとか)
	//MiddleWares []DataSaveMiddleWare `json:"middlewares"`
	// 開発言語経験
	ProgrammingLanguages []ProgrammingLanguage `json:"programminglanguages"`
	// 業務 (CIやレビュー体制でのコミットが多いやエンジニアとしてはアプリ重視かなど)
	JobDescriptions []JobDescription `json:"jobdescriptions"`
}

// 生みだしたアイテムの効果 素材や当時のスキルで効果が違う
type ItemEffect struct {
	EffectType  int    `json:"effect_type"`
	EffectValue int    `json:"effect_value"`
	CreatedAt   string `json:"created_at"`
}

type MyItem struct {
	No     int        `json:"no"`
	Name   string     `json:"name"`
	Effect ItemEffect `json:"effect"`
}

// 制作物や自伝や所有物といった自由に持ち出せるアイテム
type MyItemnList struct {
	Items []MyItem `json:"items"`
}

// 解放ミッション一覧
type MissionData struct {
	// 開発やスキル系の判定 自分の持っているアイテムで開放値
	ItemBorder int `json:"item_border"`
	// 経験値の判定

}

// ダウンロード時のデータ取得
type CarrerResult struct {
	// 各プロジェクトでのコミット時間やジャンルの一覧
	Elements []TaskElement `json:"elements"`
}

// タスクをした時のジョブ内容を表す
type TaskElement struct {
	Id             int    `json:"id"`
	ProjectGroupId int    `json:"project_group_id"`
	Name           string `json:"name"`
	JobId          int    `json:"job_id"`
	GenreId        int    `json:"genre_id"`
	ProgrammingId  int    `json:"programming_id"`
	// プロジェクトで取り扱ったサーバーの種類id
	ServerId int `json:"server_id"`
	// DB自体のインストールしているid
	DbId        int `json:"db_id"`
	TechPardId1 int `json:"tech_part_id_1"`
	TechPardId2 int `json:"tech_part_id_2"`
	// プロジェクトジョイン時の作業開始時刻
	StartAt string `json:"startat"`

	// 開発及び関連作業工数 作業時間
	Hour int `json:"hour"`

	// 経験年数表示用
	DisplayYear int `json:"display_year"`
	// 経験月数
	DisplayMonth int `json:"display_month"`

	DocumentHour int    `json:"document_hour"`
	JobLevel     int    `json:"job_level"`
	PartNote     string `json:"note"`
}

// 個人が自分のソフトごとで「自己紹介」として利用するデータ
type PersonalConvert struct {
	TaskElements []TaskElement `json:"task_elements"`
}

// 無駄開発大好きなのは、ばれてますよ

// 未来のスキル経歴のマスターデータと実行結果出力
func main() {
	// SF面接用のデータ一覧を用意
	println("contents database create start")

	startTime := time.Now()

	// 会社情報と人間情報なしでの能力一覧

	//
	jobContentsDatabase := createJocContentsDatabase()

	body, err := json.Marshal(jobContentsDatabase)
	if err != nil {
		println("parse error")
		log.Fatal(err)
		return
	}

	println("personal database outputs")

	//元データが取得できたので、書き込む配信データを開く
	f, err := os.OpenFile("masterdata/contents.json", os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModeAppend.Perm())
	if err != nil {
		println(err.Error())
		return
	}

	println("取得した要素一覧を取得")

	_, err = f.Write(body)
	if err != nil {
		println(err.Error())
		log.Fatal(err)
		return
	}

	defer f.Close()

	// 郵便番号からの履歴書生成と郵送をする
	csvF, err := os.Open("masterdata/ken_all.csv")
	if err != nil {
		println(err.Error())
		return
	}

	defer csvF.Close()

	// 郵便局設定
	scanner := bufio.NewScanner(transform.NewReader(csvF, japanese.ShiftJIS.NewDecoder()))

	var locaionInformations = map[string]*LocationInformation{}

	for scanner.Scan() {
		splitDataRow := strings.Split(scanner.Text(), ",")
		areaCode, _ := strconv.Atoi(splitDataRow[0])
		oldPostNum, _ := strconv.Atoi(splitDataRow[1])
		// 郵便局での一覧を設定する
		if splitDataRow[2] == "1520023" {
			println(splitDataRow[2])
		}
		locaionInformations[splitDataRow[2]] = &LocationInformation{
			AreaCode:       areaCode,
			OldPostNum:     oldPostNum,
			PostNum:        splitDataRow[2],
			PrefectureKana: splitDataRow[3],
			CityNameKana:   splitDataRow[4],
			TownNameKana:   splitDataRow[5],
			PrefectureName: splitDataRow[6],
			CityName:       splitDataRow[7],
			TownName:       splitDataRow[8],
		}
	}

	// データ保存処理を開始(クライアントアプリ向けに整形をする)

	// 自分自身の経歴データ
	personalf, err := os.OpenFile("personaldata/public.json", os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModeAppend.Perm())
	if err != nil {
		println(err.Error())
		return
	}

	// 一般向けデータを出力して保存
	personalData := createPersonalConvert()

	personalBasicData := &PersonalBasicInformation{
		NickName:     "かいはつのもり　じゅうにんくん",
		FromName:     "ためされるだいち",
		RealLocation: "1520023",
	}

	println("履歴書")

	println("")
	println("")

	println("基本情報")
	println("------------------------------")

	println(fmt.Sprintf("登録名称 %s", personalBasicData.NickName))
	println("------------------------------")

	println("")
	println("")

	println(fmt.Sprintf("紹介者向け出身地 %s", personalBasicData.FromName))

	println("------------------------------")

	println("")
	println("")

	location, _ := locaionInformations[fmt.Sprintf(`"%s"`, personalBasicData.RealLocation)]

	println("")
	println("")
	println("------------------------------")

	println(fmt.Sprintf("郵便番号 〒%s", location.PostNum))
	println(fmt.Sprintf("連絡先住所 %s %s %s ", location.PrefectureName, location.CityName, location.TownName))

	println("")
	println("")
	println("------------------------------")

	println("")
	println("")
	println("------------------------------")

	for i := range personalData.TaskElements {

		hour := personalData.TaskElements[i].Hour
		personalData.TaskElements[i].DisplayMonth = 0
		// １ヵ月以上の経験になった
		if hour >= MonthHour && hour >= YearHour {
			// 〇ヵ月の経験を設定
			personalData.TaskElements[i].DisplayMonth = hour % 12
		} else if hour >= MonthHour {
			// 1年未満数ヵ月の経験者
			personalData.TaskElements[i].DisplayMonth = hour / MonthHour
		}
		// 〇年以上の計算(通年雇用)
		personalData.TaskElements[i].DisplayYear = 0
		// １年以上の経験者を出力できるように設定
		if hour >= YearHour {
			personalData.TaskElements[i].DisplayYear = hour / YearHour
		}
		message := fmt.Sprintf("参加したおしごと %s 参画工数 %d  ", personalData.TaskElements[i].Name, personalData.TaskElements[i].Hour)
		println(message)
		displayMessage := fmt.Sprintf("業務内容 %s 開始日時 %s ～ %d年 %dヵ月 経験  ", personalData.TaskElements[i].PartNote, personalData.TaskElements[i].StartAt, personalData.TaskElements[i].DisplayYear, personalData.TaskElements[i].DisplayMonth)
		println(displayMessage)

	}

	body, err = json.Marshal(personalData)
	if err != nil {
		println("parse error")
		log.Fatal(err)
		return
	}

	_, err = personalf.Write(body)
	if err != nil {
		println(err.Error())
		log.Fatal(err)
		return
	}

	defer personalf.Close()

	// アプリ内で使う配信データの生成時間を出力(国や地域拡大時の改善指標と運営レポートで利用)

	// 番号と実行時間によって「開発におけるリードタイム改善」用出力
	// ミリ秒と秒単位での改善レポート
	developDataFormat := fmt.Sprintf("%d,%v", 1, time.Since(startTime).Milliseconds())

	println(developDataFormat)
	// 別kpiバッチ予定 実施した作業から経験年数とかいろいろなもの計算後の時間
	records := [][]string{
		{"No", "ExecTime"},
		{"2", strconv.FormatInt(time.Since(startTime).Milliseconds(), 10)},
	}

	// 日付でのKPIローテーション
	csvfile, err := os.Create(fmt.Sprintf("./kpi/jobparameter_exec_time_%s.csv", time.Now().UTC().Format("2006-01-02")))

	w := csv.NewWriter(csvfile)

	w.WriteAll(records)

	w.Flush()
	csvfile.Close()

}

// #推しの一コマ クイズの開設メモするんで「耐性ありそうな人」でノート作ってます
// #推しの一コマ あぁ、独身でもいいんですが「解説文作成の耐性者」一人用意してくださいｗ
// #推しの一コマ あ、もう一人の推しから「カロリーグラフでみんな脱落しました...」と連絡が来てしまいました...
func createJocContentsDatabase() *JocContentsDatabase {
	// 挨拶先建物一覧
	taskProjects := []TaskProject{
		{
			Id:          1,
			Name:        "業務アプリケーション開発",
			Description: "IOS端末への切り替え作業",
		},
		{
			Id:          2,
			Name:        "スタートアッププロジェクト",
			Description: "ベンチャーキャピタル出資による各種プロジェクト",
		},
		{
			Id:          3,
			Name:        "公共事業系Sier事務所",
			Description: "官公庁及び各省庁関連の業務請負Sier様事務所(同事務所に入ってくる各種業務担当)",
		},
		{
			Id:          4,
			Name:        "バックエンド専門会社業務",
			Description: "各ゲームのバックエンドサーバー担当会社常駐での作業",
		},
		{
			Id:          5,
			Name:        "メディア会社案件",
			Description: "既存サーバーのリプレイスプロジェクト",
		},
		{
			Id:          5,
			Name:        "メディア会社案件",
			Description: "既存サーバーのリプレイスプロジェクト",
		},
		{
			Id:          6,
			Name:        "自治体会計システム刷新",
			Description: "セキュリティパッチ発生時におけるフレームワーク切り替えのサポート業務",
		},
		{
			Id:          7,
			Name:        "大手メガバンクの詳細設計～詳細テスト",
			Description: "外部インターフェースを確認と、各種設計書の続き",
		},
		{
			Id:          15,
			Name:        "総合EC及び一部ビューワーの本開発フェーズのプログラミング~支援",
			Description: "グループ会社でのSDK及びサーバー切り替え時の新規システム開発支援",
		},
	}

	/*	genres := []ProductGenre{
		{
			Id:   1,
			Name: "ブラウザゲーム",
		},
		{
			Id:   2,
			Name: "財務・会計",
		},
		{
			Id:   3,
			Name: "Webアプリ",
		},
		{
			Id:   4,
			Name: "メディア関連",
		},
		{
			Id:   5,
			Name: "スマホゲーム",
		},
		{
			Id:   6,
			Name: "スマホネイティブアプリ",
		},
		{
			Id:   7,
			Name: "業務アプリケーション",
		},
		{
			Id:   8,
			Name: "社内業務基幹",
		},
		{
			Id:   9,
			Name: "位置情報ゲームと管理システム",
		},
		{
			Id:   10,
			Name: "ECサイト(購入まで、別途ビューワーがあるんものは時間分割)",
		},
		{
			Id:   11,
			Name: "PCオンラインゲーム",
		},
		{
			Id:   12,
			Name: "会員向けマルチデバイス対応アプリ(スマホ、Web、PCなどの他展開関連はここ)",
		},
		{
			Id:   13,
			Name: "メールシステム",
		},
	}*/

	/*	benders := []BenderCategory{
			{
				Id:         1,
				BenderType: 1,
				Name:       "aws",
			},
			{
				Id:         2,
				BenderType: 1,
				Name:       "gcp",
			},
			{
				Id:         3,
				BenderType: 1,
				Name:       "aws",
			},
			{
				Id:         4,
				BenderType: 1,
				Name:       "オンプレ",
			},
			{
				Id:         5,
				BenderType: 1,
				Name:       "さくらクラウド",
			},
			{
				Id:         6,
				BenderType: 2,
				Name:       "JP1(日本語ではなく、固有ミドルウェア)",
			},
			{
				Id:         7,
				BenderType: 1,
				Name:       "gmo",
			},
		}

		middleWares := []DataSaveMiddleWare{
			{
				Id:       1,
				Category: 1,
				Name:     "mysql",
			},
			{
				Id:       2,
				Category: 1,
				Name:     "oracle",
			},
			{
				Id:       3,
				Category: 1,
				Name:     "datastore",
			},
			{
				Id:       4,
				Category: 1,
				Name:     "spanner",
			},
			{
				Id:       5,
				Category: 1,
				Name:     "mariadb",
			},
			{
				Id:       6,
				Category: 2,
				Name:     "redis",
			},
		}*/

	languagues := []ProgrammingLanguage{
		{
			Id:       1,
			Name:     "PHP",
			TotalSum: 6000,
		},
		{
			Id:       2,
			Name:     "JavaScript",
			TotalSum: 180,
		},
		{
			Id:       3,
			Name:     "Perl",
			TotalSum: 300,
		},
		{
			Id:       4,
			Name:     "Java",
			TotalSum: 600,
		},
		{
			Id:       5,
			Name:     "Python",
			TotalSum: 5000,
		},
		{
			Id:       6,
			Name:     "Go",
			TotalSum: 4500,
		},
		{
			Id:       7,
			Name:     "C#(Unity)",
			TotalSum: 0,
		},
		{
			Id:       8,
			Name:     "SQL+α",
			TotalSum: 1000,
		},
		{
			Id:       9,
			Name:     "ActionScript2.0",
			TotalSum: 200,
		},
		{
			Id:       10,
			Name:     "Ruby",
			TotalSum: 500,
		},
		{
			Id:       11,
			Name:     "Solidity",
			TotalSum: 200,
		},
		{
			Id:       12,
			Name:     "C#(Server)",
			TotalSum: 1500,
		},
		{
			Id:       13,
			Name:     "C++",
			TotalSum: 0,
		},
		{
			Id:       14,
			Name:     "他(インフラや別パート手伝い)",
			TotalSum: 0,
		},
		{
			Id:       15,
			Name:     "VBA",
			TotalSum: 0,
		},
		{
			Id:       16,
			Name:     "Shell",
			TotalSum: 0,
		},
		{
			Id:       17,
			Name:     "OrganizeAndProjectProfit(DataAnalizeQuery)",
			TotalSum: 50,
		},
	}
	jobDescriptions := []JobDescription{
		JobDescription{
			Id:       1,
			Category: 1,
			Name:     "コードレビュー(CodeReview)",
		},
		JobDescription{
			Id:       2,
			Category: 1,
			Name:     "DB設計(DataBaseSchemaDocument)",
		},
		JobDescription{
			Id:       3,
			Category: 1,
			Name:     "テストコード(UnitProgramTest)",
		},
		JobDescription{
			Id:       4,
			Category: 1,
			Name:     "インスタンス設計(AppSeverRequestScenarioDocument)",
		},
		JobDescription{
			Id:       5,
			Category: 1,
			Name:     "アプリ運営前フロー(総合試験及びリリースデータ用)",
		},
	}

	return &JocContentsDatabase{
		TaskProjects: taskProjects,
		//	ProductGenres:        genres,
		//	Benders:              benders,
		//	MiddleWares:          middleWares,
		ProgrammingLanguages: languagues,
		JobDescriptions:      jobDescriptions,
	}
}

func createPersonalConvert() *PersonalConvert {
	return &PersonalConvert{
		TaskElements: []TaskElement{
			//
			TaskElement{
				Id:             1,
				ProjectGroupId: 2,
				Name:           "業務アプリケーションにおけるフロントエンド",
				JobId:          7,
				GenreId:        2,
				ProgrammingId:  4,
				// プロジェクトで取り扱ったサーバーの種類id
				ServerId: 1,
				// DB自体のインストールしているid
				DbId: 2,
				// プロジェクトジョイン時の作業開始時刻
				StartAt: "2010.11",

				// 開発及び関連作業工数
				Hour:         80,
				DocumentHour: 0,
				JobLevel:     1,
				PartNote:     "開発",
			},
			TaskElement{
				Id:             2,
				ProjectGroupId: 2,
				Name:           "ブラウザゲーム開発",
				JobId:          2,
				GenreId:        9,
				ProgrammingId:  4,
				// プロジェクトで取り扱ったサーバーの種類id
				ServerId: 1,
				// DB自体のインストールしているid
				DbId: 3,
				// プロジェクトジョイン時の作業開始時刻
				StartAt: "2011.13",

				// 開発及び関連作業工数
				Hour:         80,
				DocumentHour: 80,
				JobLevel:     1,
				PartNote:     "詳細設計～製造",
			},
			TaskElement{
				Id:             4,
				ProjectGroupId: 4,
				Name:           "ブラウザゲーム開発",
				JobId:          2,
				GenreId:        1,
				ProgrammingId:  4,
				// プロジェクトで取り扱ったサーバーの種類id
				ServerId: 1,
				// DB自体のインストールしているid
				DbId: 1,
				// プロジェクトジョイン時の作業開始時刻
				StartAt: "2011.11",
				// 開発及び関連作業工数
				Hour:         1600,
				DocumentHour: 0,
				JobLevel:     1,
				PartNote:     "運営ツール設計実装",
			},
		},
	}
}
