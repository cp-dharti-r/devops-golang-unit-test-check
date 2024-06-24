package test

import "github.com/jmoiron/sqlx"

func TestDB() *sqlx.DB {
	db := sqlx.MustConnect("mysql", "root:password@tcp(localhost:3306)/my_test_db")
	
	defer db.Close()
	
	return db
}
