package products_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// 商品エンティティ、値オブジェクトのテスト用エントリーポイント
func TestEntityPackage(t *testing.T) {
	RegisterFailHandler(Fail)                      // テストが失敗した場合のハンドラーを登録
	RunSpecs(t, "domain/models/productsパッケージのテスト") // 全てのテストを実行
}
