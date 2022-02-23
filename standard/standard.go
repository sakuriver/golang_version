package standard

// コンテンツ内の登場キャラクター
type ContentsCharacter struct {
	Id   int
	Name string
}

// 転生系
func GetParties() []*ContentsCharacter {
	contentsCharacters := []*ContentsCharacter{}

	// 転生もので一緒に登録されているメンバー
	contentsCharacters = append(contentsCharacters, &ContentsCharacter{
		Id:   200,
		Name: "カイ",
	})
	contentsCharacters = append(contentsCharacters, &ContentsCharacter{
		Id:   300,
		Name: "John",
	})
	return contentsCharacters
}
