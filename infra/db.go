package infra

import (
	"fmt"
	"gin-todo-app/models"
	"log"
	"os"

	"github.com/jaevor/go-nanoid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	nanoIDCharacters = "abcdefghijklmnopqrstuvwxyz"
	nanoIDLength     = 16
)

func SetupDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=Asia/Tokyo",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)

	var (
		db  *gorm.DB
		err error
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	log.Println("Connection Opened to Database")

	return db
}

func TestDB() *gorm.DB {
	dbName := fmt.Sprintf("%s_%s_test", os.Getenv("DB_NAME"), generateRandomString())
	if err := createPosgresDatabase(dbName); err != nil {
		panic("failed to connect database")
	}
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=Asia/Tokyo",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		dbName,
		os.Getenv("DB_PASSWORD"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	if err := db.AutoMigrate(&models.Item{}, &models.User{}, &models.Status{}); err != nil {
		panic("Failed to migrate database")
	}
	return db
}

func createPosgresDatabase(dbName string) (err error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=Asia/Tokyo",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// テスト用のDBを作成
	createDatabaseCommand := fmt.Sprintf("CREATE DATABASE %s", dbName)
	tx := db.Exec(createDatabaseCommand)
	if tx.Error != nil {
		return tx.Error
	}

	enrolUserCommand := fmt.Sprintf("GRANT ALL PRIVILEGES ON DATABASE %s TO %s", dbName, os.Getenv("DB_USER"))
	// テストで使うユーザーに権限を付与
	tx = db.Exec(enrolUserCommand)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func generateRandomString() string {
	return must(nanoid.CustomASCII(nanoIDCharacters, nanoIDLength))()
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
