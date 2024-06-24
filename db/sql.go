package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx")

func NewSql() *sqlx.DB {
	db := sqlx.MustConnect("mysql", "root:password@tcp(localhost:3306)/my_db")
	
	defer db.Close()
	
	return db
}
