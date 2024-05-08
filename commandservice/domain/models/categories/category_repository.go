package categories

import (
	"context"
	"database/sql"
)

type CategoryRepository interface {
	// 同名の商品カテゴリが存在するか確認する
	Exists(ctx context.Context, tran *sql.Tx, category *Category) error
	// あたらしい商品カテゴリを登録する
	Create(ctx context.Context, tran *sql.Tx, category *Category) error
	// 商品カテゴリを更新する
	UpdateById(ctx context.Context, tran *sql.Tx, category *Category) error
	// 商品カテゴリを削除する
	DeleteById(ctx context.Context, tran *sql.Tx, category *Category) error
}
