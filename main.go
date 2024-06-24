package main

import (
	"db"
	"user"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var sqlDb *sqlx.DB

func init() {
	sqlDb = db.NewSql()
}

func main() {
	router := gin.Default()

	userRepo := user.New(sqlDb)

	router.POST("/api/users", userRepo.Create)
	router.GET("/api/users/:id", userRepo.Get)
	router.PUT("/api/users/:id", userRepo.Update)
	router.DELETE("/api/users/:id", userRepo.Delete)

	router.Run(":8000")

	defer sqlDb.Close()
}
