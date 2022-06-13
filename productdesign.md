今回作るもの

既知のフォーマットとして、以下のページを利用しています。
https://www.altexsoft.com/blog/business/technical-documentation-in-software-development-types-best-practices-and-tools/

# summary

bffや各種マルチサーバーにより「機能全体で一つの空間」を「共有やデータ連動あり」でも進められるようになったと考えている。

技術再習得については、以下を兼ねている。

機能をお店ととらえて「リアル店舗の開店やアップデート」
→マイクロトランザクションサーバー

各アプリ向けの共通基盤サーバーとレスポンス用サーバー
→BFF及びサーバー間通信


default standard format 
    bff etc request and response standard format prototype 
        xr and digitaltwin app library
        multi app data share 

    techinical learn

        bff server               ... multi device gear server
        mictrotransaction server ... xr or tool function base server

## Requirements(構成要素)

将来的なマルチデバイスサーバーを兼ねている
最大で、以下のようなシナリオを想定(max scenario)

multi device and after covid19 bussiness full
shoppinng and daily 3d place

VRのお店→特定のアバターと会話して購入
max scenario

vr store login → buyer avater talk → purchase

### User Story Title(アプリを使ってもらう時の流れ)

1. お店に入って商品を購入(store login puroduct purchased scenario)
アプリ起動 → 販売店に入る → 販売中の売り子idを指定 → 商品を確認 → 購入 → 商品を確認

app launch → store in → buyer avater talk → product select → purchase → my purchased item list show

2.  購入した商品の使用
アプリ起動 → 自分のカバンを選択 → 所持アイテム一覧を選択 → アイテムの試着、試飲、試用


### User Story Map(開発及び各種の流れ)

#### User Tasks 


##### Release1
    ・スマホアプリ
    ・PCクライアントアプリ


##### Release2



##### Release3



## UserInterface and design


　