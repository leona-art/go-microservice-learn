package handler

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type DBCofig struct {
	Dbname string `toml:"dbname"`
	Host   string `toml:"host"`
	Port   int    `toml:"port"`
	User   string `toml:"user"`
	Pass   string `toml:"pass"`
}

func tomlRead() (*DBCofig, error) {
	// 環境変数からパスを取得
	path := os.Getenv(("DATABASE_TOML_PATH"))
	if path == "" {
		path = "infra/sqlboiler/handler/database.toml"
	}
	m := map[string]DBCofig{}
	_, err := toml.DecodeFile(path, &m)
	if err != nil {
		return nil, err
	}

	config := m["mysql"]

	return &config, nil
}

func DbConnect() error {
	config, err := tomlRead()
	if err != nil {
		return DBErrHandler(err)
	}

	// 接続文字列の生成
	rdbms := "mysql"
	connect_str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Pass, config.Host, config.Port, config.Dbname)
	conn, err := sql.Open(rdbms, connect_str)
	if err != nil {
		return DBErrHandler(err)
	}

	// データベース接続を確認する
	if err := conn.Ping(); err != nil {
		return DBErrHandler(err)
	}

	MAX_IDLE_CONNS := 10
	MAX_OPEN_CONNS := 100
	CONN_MAX_LIFETIME := 300 * time.Second

	conn.SetMaxIdleConns(MAX_IDLE_CONNS)
	conn.SetMaxOpenConns(MAX_OPEN_CONNS)
	conn.SetConnMaxLifetime(CONN_MAX_LIFETIME)
	boil.SetDB(conn)
	boil.DebugMode = true
	return nil
}
