package dbClient

import (
	"database/sql"
	"fmt"
	"garg/constants"
	db "garg/db/generated"
	"garg/resources"
	"path/filepath"

	_ "modernc.org/sqlite"
	"github.com/pressly/goose/v3"
)

func Migrate() {
	if err := goose.SetDialect("sqlite"); err != nil {
		fmt.Println(err)
	}

	sqlDB, err := goose.OpenDBWithDriver("sqlite", filepath.Join(constants.DataDir, "garg.db"))
	if err != nil {
		fmt.Println(err)
	}

	goose.SetBaseFS(resources.Resources)

	if err := goose.Up(sqlDB, "db_migrations"); err != nil {
		fmt.Println(err)
	}
}

func GetClient() (*db.Queries, error) {
	conn, err := sql.Open("sqlite", filepath.Join(constants.DataDir, "garg.db"))
	if err != nil {
		return nil, err
	}

	client := db.New(conn)
	return client, nil
}
