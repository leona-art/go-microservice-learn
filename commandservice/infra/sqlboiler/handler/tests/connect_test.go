package handler_test

import (
	"commandservice/infra/sqlboiler/handler"
	"os"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestConn(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "infra/sqlboiler/handlerパッケージのテスト")
}

var _ = Describe("データベース接続テスト", func() {
	It("接続が成功した場合,nilが帰る", Label("DB接続"), func() {
		absPath, _ := filepath.Abs("../../database.toml")

		os.Setenv("DATABASE_TOML_PATH", absPath)
		result := handler.DbConnect()
		Expect(result).To(BeNil())
	})
})
