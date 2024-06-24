package main

import "github.com/jmoiron/sqlx"

func TestDB() *sqlx.DB {
	db := sqlx.MustConnect("mysql", "root:password@tcp(localhost:3306)/test_db")
	return db
}
