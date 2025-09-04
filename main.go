package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"go-fram/config"
	"go-fram/internal"
)

func main() {
	dialect := os.Getenv("DB_DIALECT")
	if dialect == "" {
		dialect = "sqlite"
	}
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "auth.db"
	}
	db, err := config.ConnectDB(dialect, dsn)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&internal.User{})

	r := gin.Default()
	r.POST("/register", internal.Register(db))
	r.POST("/login", internal.Login(db))
	r.GET("/protected", internal.AuthMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "You are authenticated!"})
	})

	r.Run(":8080")
}
