package db

import (
	"config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewSql() *sqlx.DB {
	cfg := config.LoadConfig()

	dsn := cfg.DBUser + ":" + cfg.DBPassword + "@tcp(" + cfg.DBHost + ":" + cfg.DBPort + ")/" + cfg.DBName
	db := sqlx.MustConnect("mysql", dsn)

	return db
}
