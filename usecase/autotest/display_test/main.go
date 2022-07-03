package main

import (
	"fmt"
	"time"

	"github.com/sclevine/agouti"
)

var domainRoot = "https://gcp-auto-test-viewer.web.app/"

// 自動テスト
// 想定通りに動いているかをボタンで確認をする

var screenShopTempPath = ""
var screenShopTempStartPath = ""

var inputAnimationBasePath = ""

func main() {
	// Chromeを利用することを宣言
	options := agouti.ChromeOptions(
		"args", []string{

			"--disable-logging",
		})
	excludeSwitches := agouti.ChromeOptions(
		"excludeSwitches", []string{
			"enable-logging",
		})
	agoutiDriver := agouti.ChromeDriver(options, excludeSwitches)
	defer agoutiDriver.Stop()
	err := agoutiDriver.Start()
	if err != nil {
		panic(err)
	}

	page, err := agoutiDriver.NewPage()

	if err != nil {
		panic(err.Error())
	}

	// エラーテスト確認のwebアプリを開く
	page.Navigate(domainRoot)
	submitButton := page.FindByButton("submitButton")

	// 最初エラー
	submitButton.Click()
	time.Sleep(3 * time.Second)
	submitButton.Submit()

	page.Screenshot(screenShopTempStartPath)
	time.Sleep(1 * time.Second)

	searchText := page.FindByName("search")
	messages := []string{"こ", "ん", "にちは"}

	messageText := ""
	for i := range messages {
		messageText += messages[i]
		searchText.Fill(messageText)
		time.Sleep(1 * time.Second)
		page.Screenshot(fmt.Sprintf(inputAnimationBasePath, i))
	}

	// ページを開いた後に、スクショを撮影
	page.Screenshot(screenShopTempPath)

	// 2回目通過用
	submitButton.Submit()

}
