package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//ConnectDB connects go to mysql database
func ConnectDB() *gorm.DB {

	env := "backend.env"

	dir := "/var/www/boxinvesting.ru/"
	environmentPath := filepath.Join(dir + env)
	fmt.Println(environmentPath)

	errorENV := godotenv.Load(environmentPath)

	fmt.Println(dir + env)
	if errorENV != nil {
		fmt.Println(errorENV)
		panic("Failed to load env file ConnectDB")
		log.Fatal(errorENV)
	}
	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//if err != nil {
	//	log.Fatal(err)
	//	}
	//fmt.Println(dir)
	//	environmentPath := filepath.Join(dir, ".env")
	//	err = godotenv.Load(environmentPath)
	//	panic(err)
	//	fmt.Println(err)
	//errorENV := godotenv.Load()
	//errorENV := godotenv.Load(filepath.Join(path_dir, ".env"))

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	fmt.Println(dbUser)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=false&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, errorDB := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errorDB != nil {
		panic("Failed to connect mysql database")
	}

	return db
}

//DisconnectDB is stopping your connection to mysql database
func DisconnectDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to kill connection from database")
	}
	dbSQL.Close()
}
