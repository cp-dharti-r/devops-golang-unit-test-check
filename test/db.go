package test

import (
	"config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func TestDB() *sqlx.DB {
	cfg := config.LoadConfig()

	dsn := cfg.DBUser + ":" + cfg.DBPassword + "@tcp(" + cfg.DBHost + ":" + cfg.DBPort + ")/my_test_db"
	db := sqlx.MustConnect("mysql", dsn)

	return db
}
