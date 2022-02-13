package main

import (
	"fmt"
	"testing"
)

// 商品の購入テスト
func TestPurchaseSuccess(t *testing.T) {

	// 商品を1個指定して購入を実施する
	pinfo := PlayerInformation{}

	pinfo.Pid = 1
	pinfo.Gold = 10000
	pinfo.Playeritems = map[int]int{}

	// アイテムを1個選択
	psp := PlayerShopPurchase{}
	psp.Shopid = 2
	psp.Num = 1

	// 購入開始ログを事前に残しておく
	purchaseRequestMessage := fmt.Sprintf("purchase start shopid %[1]d selectnum %[2]d", psp.Shopid, psp.Num)
	t.Log(purchaseRequestMessage)

	var purcharseResult = Purchase(pinfo, psp)

	// 購入処理が完了後のプレイヤー情報がコピーされている
	if pinfo.Pid != purcharseResult.Pid {
		t.Errorf("purchasefailed")
	}

	// 購入後に金額が増えていたらエラー
	if pinfo.Gold < purcharseResult.Gold {
		t.Errorf("gold subtract error %d、 %d", pinfo.Gold, purcharseResult.Gold)

	}

	// 購入後に商品が増えていなかったらエラー
	pitemnum, pitemexists := pinfo.Playeritems[psp.Shopid]

	if pitemexists == false {
		t.Errorf("item purchase error")
	}

	purchaseMessage := fmt.Sprintf("item purchase after itemid %[1]d num %[2]d", psp.Shopid, pitemnum)
	resultMessage := fmt.Sprintf("player item purchase gold %[1]d ", purcharseResult.Gold)
	// 商品購入結果をログとして出力
	t.Log(purchaseMessage)
	t.Log(resultMessage)

}
