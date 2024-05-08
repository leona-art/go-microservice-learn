package products

import (
	"commandservice/errs"
	"fmt"
)

// 商品単価を表す値オブジェクト
type ProductPrice struct {
	value uint32
}

// 値のゲッター
func (ins *ProductPrice) Value() uint32 {
	return ins.value
}

// 同一性の確認
func (ins *ProductPrice) Equals(other *ProductPrice) bool {
	return ins.value == other.Value()
}

// コンストラクタ
func NewProductPrice(value uint32) (*ProductPrice, *errs.DomainError) {
	const MIN_VALUE uint32 = 50
	const MAX_VALUE uint32 = 100000

	if value < MIN_VALUE || value > MAX_VALUE {
		return nil, errs.NewDomainError(
			fmt.Sprintf("商品単価は%d以上%d以下でなければなりません", MIN_VALUE, MAX_VALUE),
		)
	}
	return &ProductPrice{value: value}, nil
}
