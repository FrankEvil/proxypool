package database

import (
	"fmt"
	"os"

	"github.com/zu1k/proxypool/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connect() (err error) {
	dsn := "root:frankevil0821@tcp(frankevil.tk:3306)/proxypool?charset=utf8mb4&parseTime=True&loc=Local"
	if url := config.Config.DatabaseUrl; url != "" {
		dsn = url
	}
	if url := os.Getenv("DATABASE_URL"); url != "" {
		dsn = url
	}
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err == nil {
		fmt.Println("DB connect success: ", DB.Name())
	}
	return
}
