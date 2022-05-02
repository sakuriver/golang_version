package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type ShopBuyerProduct struct {
	Id           string `json:"id"`
	ShopMasterId string `json:"shop_master_id"`
	BuyerId      string `json:"buyer_id"`
	PackageUrl   string `json:"package_url"`
	Num          int    `json:"num"`
	Rate         string `json:"rate"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

func main() {
	var inTE, outTE *walk.TextEdit

	var tv *walk.TableView

	MainWindow{
		Title:   "SCREAMO",
		MinSize: Size{600, 400},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			PushButton{
				Text: "SCREAM",
				OnClicked: func() {

					client := &http.Client{}

					// api net work
					// --header 'Content-Type: application/json' \
					req, _ := http.NewRequest("GET", "https://stoplight.io/mocks/productchallengeweb/shop-mock-server/43861283/products/buyerid", nil)
					req.Header.Add("Content-Type", "application/json")
					req.Header.Add("Prefer", "code=200, example=example-1")
					resp, _ := client.Do(req)
					body, _ := io.ReadAll(resp.Body)

					var products []ShopBuyerProduct

					json.Unmarshal(body, &products)
					var resultStr = inTE.Text()
					for _, v := range products {
						resultStr += fmt.Sprintf("id:%s\r\n", v.Id)
						resultStr += fmt.Sprintf("shop_master_id:%s\r\n", v.ShopMasterId)
						resultStr += fmt.Sprintf("package_url:%s\r\n", v.PackageUrl)
						resultStr += fmt.Sprintf("num:%d\r\n", v.Num)
						resultStr += fmt.Sprintf("rate:%s\r\n", v.BuyerId)
						resultStr += fmt.Sprintf("created_at:%s\r\n", v.CreatedAt)
						resultStr += fmt.Sprintf("updated_at:%s\r\n", v.UpdatedAt)
					}

					outTE.SetText(strings.ToUpper(resultStr))

				},
			},
			TableView{
				AssignTo:         &tv,
				AlternatingRowBG: true,
				CheckBoxes:       true,
				ColumnsOrderable: true,
				MultiSelection:   true,
				Columns: []TableViewColumn{
					{Title: "#"},
					{Title: "ショップマスター番号"},
					{Title: "数量"},
				},
				StyleCell: func(style *walk.CellStyle) {
				},
			},
		},
	}.Run()
}
