package conf_db

import (
	"fmt"
	"github.com/gandra/my-bookmarks/database/db"
	"github.com/gandra/my-bookmarks/domain/bookmarks"
	"github.com/gandra/my-bookmarks/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"
	"os"
)

const (
	db_users_username = "db_users_username"
	db_users_password = "db_users_password"
	db_users_host     = "db_users_host"
	db_users_schema   = "db_users_schema"
)

var (
	username = os.Getenv(db_users_username)
	password = os.Getenv(db_users_password)
	host     = os.Getenv(db_users_host)
	schema   = os.Getenv(db_users_schema)
)

func InitDatabase() {
	// TOD: Fix temp hack to avoid playing with env vars
	username = "postgres"
	password = "mysecret"
	host = "localhost"
	schema = "bookmarks"

	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432", host, username, password, schema)
	config := gorm.Config{
		Logger: logger2.Default.LogMode(logger2.Warn),
	}
	db.Client, err = gorm.Open(postgres.Open(dsn), &config)
	if err != nil {
		fmt.Println(err)
		panic("Failed to connect to database!")
	}
	logger.Info("Database connected!")
	db.Client.AutoMigrate(&bookmarks.Bookmark{})
	logger.Info("Migrated DB")

}
