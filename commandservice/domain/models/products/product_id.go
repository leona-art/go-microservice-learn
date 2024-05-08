package products

import (
	"commandservice/errs"
	"fmt"
	"regexp"
	"unicode/utf8"
)

// 商品番号を表す値オブジェクト
type ProductId struct {
	value string
}

// 値のゲッター
func (ins *ProductId) Value() string {
	return ins.value
}

// 同一性の確認
func (ins *ProductId) Equals(other *ProductId) bool {
	return ins.value == other.Value()
}

// コンストラクタ
func NewProductId(value string) (*ProductId, *errs.DomainError) {
	// フィールドの長さ
	const LENGTH int = 36

	// UUIDの正規表現
	const REGEXP string = `^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`

	// 引数の文字数チェック
	if utf8.RuneCountInString(value) != LENGTH {
		return nil, errs.NewDomainError(
			fmt.Sprintf("商品番号は%d文字でなければなりません", LENGTH),
		)

	}

	// 引数の正規表現チェック
	if !regexp.MustCompile(REGEXP).MatchString(value) {
		return nil, errs.NewDomainError(
			"商品番号はUUID形式でなければなりません",
		)
	}
	return &ProductId{value: value}, nil

}
