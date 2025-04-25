module test

go 1.21.3

require (
	config v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.8.1
	github.com/jmoiron/sqlx v1.4.0
)

require filippo.io/edwards25519 v1.1.0 // indirect

replace config => ../config
