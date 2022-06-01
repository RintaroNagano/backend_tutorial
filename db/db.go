package db

import(
	"os"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func GormConnect() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Fail to read .env file : %v", err)
	}

	// .envを取得、代入
	DBMS := os.Getenv("DIALECT")
	USER := os.Getenv("USER_NAME")
	PASS := os.Getenv("PASSWORD")
	PROTOCOL := os.Getenv("PROTOCOL")
	DBNAME := os.Getenv("DB_NAME")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	DB = db
}