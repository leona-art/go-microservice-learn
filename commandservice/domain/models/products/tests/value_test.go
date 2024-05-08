package products_test

import (
	"commandservice/domain/models/products"
	"commandservice/errs"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Productエンティティを構成する値オブジェクト", Ordered,
	Label("ProductIdの生成"), func() {
		var empty_str *errs.DomainError    // 空文字 長さ36に違反
		var length_over *errs.DomainError  // 長さ36に違反
		var not_uuid *errs.DomainError     // UUID形式に違反
		var product_id *products.ProductId //UUID文字列を指定
		var uid string
		BeforeAll(func() {
			_, empty_str = products.NewProductId("")
			_, length_over = products.NewProductId("1234567890123456789012345678901234567")
			_, not_uuid = products.NewProductId("12345678-9012-3456-7890-123456789012")
			id, _ := uuid.NewRandom()
			uid = id.String()
			product_id, _ = products.NewProductId(id.String())
		})

		// 文字数の検証
		Context("文字数の検証", func() {
			It("空文字の場合", func() {
				Expect(empty_str).To(Equal(errs.NewDomainError("商品番号は36文字でなければなりません")))
			})
			It("36文字を超える場合", func() {
				Expect(length_over).To(Equal(errs.NewDomainError("商品番号は36文字でなければなりません")))
			})
		})

		// UUID形式の検証
		Context("UUID形式の検証", func() {
			It("UUID形式でない場合", func() {
				Expect(not_uuid).To(Equal(errs.NewDomainError("商品番号はUUID形式でなければなりません")))
			})
			It("UUID形式の場合", func() {
				id, _ := products.NewProductId(uid)
				Expect(product_id).To(Equal(id))
			})
		})
	})
