package main

import (
	"time"

	"github.com/sclevine/agouti"
)

var domainRoot = "https://xd.adobe.com/embed/789fab6f-4174-4b21-9979-3b9369b19f4a-2dcc/"

// 電気代XRユースケース

func main() {
	// Chromeを利用することを宣言
	agoutiDriver := agouti.ChromeDriver(agouti.Browser("chrome"))
	err := agoutiDriver.Start()
	if err != nil {
		panic(err)
	}
	defer agoutiDriver.Stop()
	page, err := agoutiDriver.NewPage()

	if err != nil {
		panic(err.Error())
	}

	// 集計対象のwebアプリを開く
	page.Navigate(domainRoot)
	time.Sleep(3 * time.Second)

	// ページを開いた後に、スクショを撮影
	page.Screenshot("openElectoricScreenShot.png")

	// 電気代の使用量ページへ遷移
	page.Navigate(domainRoot + "screen/b5cd7dfe-7af4-499d-b136-b65031474030")

	time.Sleep(2 * time.Second)

	page.Screenshot("openElectoricViewerResult.png")

	time.Sleep(2 * time.Second)

	// 使用量設定アドレス確認ページへ遷移
	page.Navigate(domainRoot + "screen/41deed8f-a791-46d8-8aee-f578c2307e8b")

	page.Screenshot("openElectoricAddConfirmResult.png")

	time.Sleep(2 * time.Second)

	// 支払い設定完了ページへの遷移
	page.Navigate(domainRoot + "screen/aac8c4d1-ebd8-4f29-abb2-7f1c029c83be")

	time.Sleep(2 * time.Second)

	page.Screenshot("openElectoricAddSettingComplete.png")

}
