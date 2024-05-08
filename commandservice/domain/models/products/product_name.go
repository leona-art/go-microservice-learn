package products

import (
	"commandservice/errs"
	"fmt"
	"unicode/utf8"
)

// 商品名を表す値オブジェクト
type ProductName struct {
	value string
}

// 値のゲッター
func (ins *ProductName) Value() string {
	return ins.value
}

// 同一性の確認
func (ins *ProductName) Equals(other *ProductName) bool {
	return ins.value == other.Value()
}

// コンストラクタ
func NewProductName(value string) (*ProductName, *errs.DomainError) {
	const MIN_LENGTH int = 5
	const MAX_LENGTH int = 30

	count := utf8.RuneCountInString(value)
	if count < MIN_LENGTH || count > MAX_LENGTH {
		return nil, errs.NewDomainError(
			fmt.Sprintf("商品名は%d文字以上%d文字以下でなければなりません", MIN_LENGTH, MAX_LENGTH),
		)
	}
	return &ProductName{value: value}, nil
}
