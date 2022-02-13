package main

// 購入処理のロジック部分単体テスト

// プレイヤーの所持金情報
type PlayerInformation struct {
	Pid         int
	Gold        int
	Playeritems map[int]int
}

// プレイヤーの注文情報
type PlayerShopPurchase struct {
	// ショップ内での番号
	Shopid int
	// こうにゅうすう
	Num int
}

// プレイヤーの所持アイテム
// 購入するとECサイト同様に自分の所持数が増える
type PlayerItem struct {
	// マスターに登録されているid
	Itemid int
	// 所持数
	num int
}

// ショップアイテム情報
type ShopItem struct {
	Id     int
	Itemid int
	Price  int
	Num    int
}

// ショップでのアイテム購入処理 で購入後の情報を返す
// 単体のイメージとして
func Purchase(pinfo PlayerInformation, psp PlayerShopPurchase) PlayerInformation {
	shopMaps := map[int]ShopItem{}

	shopFirstItem := new(ShopItem)
	shopFirstItem.Id = 1
	shopFirstItem.Itemid = 30
	shopFirstItem.Price = 200
	shopFirstItem.Num = 1

	// まとめ買い
	shopSecondItem := new(ShopItem)
	shopSecondItem.Id = 2
	shopSecondItem.Itemid = 30
	shopSecondItem.Price = 900
	shopSecondItem.Num = 5

	// テスト仕様書上がりのエンジニア向けにDBの手前として、メモリ内のアイテム購入から
	// アイテム取得処理は後でDBに差し替え
	shopMaps[shopFirstItem.Id] = *shopFirstItem
	shopMaps[shopSecondItem.Id] = *shopSecondItem

	// プレイヤーid未設定のエラー情報を返す

	shopMapResult, itemexists := shopMaps[psp.Shopid]
	if itemexists == false {
		return PlayerInformation{}
	}

	// アイテムの単価 * 選択したアイテム数
	var itemPrice = shopMapResult.Price * psp.Num

	if pinfo.Gold < itemPrice {
		return PlayerInformation{}
	}

	// お金の減算処理
	pinfo.Gold = pinfo.Gold - itemPrice
	_, shopitemexists := pinfo.Playeritems[shopMapResult.Id]

	// 商品の設定個数 * 注文数
	var shopItemValue = shopMapResult.Num * psp.Num

	if shopitemexists == false {
		// 新規購入設定
		pinfo.Playeritems[shopMapResult.Id] = shopItemValue

	} else {
		// 追加購入処理
		pinfo.Playeritems[shopMapResult.Id] += shopItemValue
	}

	// 購入してアイテム取得後の結果を設定する
	return pinfo

}
