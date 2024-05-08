package products

import (
	"commandservice/domain/models/categories"
	"commandservice/errs"

	"github.com/google/uuid"
)

// 商品エンティティ
type Product struct {
	id       *ProductId
	name     *ProductName
	price    *ProductPrice
	category *categories.Category
}

// ゲッター
func (ins *Product) Id() *ProductId {
	return ins.id
}
func (ins *Product) Name() *ProductName {
	return ins.name
}
func (ins *Product) Price() *ProductPrice {
	return ins.price
}
func (ins *Product) Category() *categories.Category {
	return ins.category
}

// 値の変更
func (ins *Product) ChangeProductName(name *ProductName) {
	ins.name = name
}
func (ins *Product) ChangeProductPrice(price *ProductPrice) {
	ins.price = price
}
func (ins *Product) ChangeProductCategory(category *categories.Category) {
	ins.category = category
}

// 同一性検証
func (ins *Product) Equals(other *Product) (bool, *errs.DomainError) {
	if other == nil {
		return false, errs.NewDomainError("比較対象がnilです")
	}
	result := ins.id.Equals(other.Id())
	return result, nil
}

// コンストラクタ
func NewProduct(name *ProductName, price *ProductPrice, category *categories.Category) (*Product, *errs.DomainError) {
	if uid, err := uuid.NewRandom(); err != nil {
		return nil, errs.NewDomainError(err.Error())
	} else if id, err := NewProductId(uid.String()); err != nil {
		return nil, err
	} else {
		return &Product{
			id:       id,
			name:     name,
			price:    price,
			category: category,
		}, nil

	}
}

// 商品エンティティの再構築

func RebuildProduct(id *ProductId, name *ProductName, price *ProductPrice, category *categories.Category) (*Product, *errs.DomainError) {
	return &Product{
		id:       id,
		name:     name,
		price:    price,
		category: category,
	}, nil
}
