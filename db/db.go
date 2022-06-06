package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

// ※ 動かないけどこうであって欲しいて言うの書いていくね
//[Motto] Void関数
func GormConnect() error {
	// .envを取得、代入
	err := godotenv.Load(".env")
	if err != nil {
		//[Good] errorハンドリング出来ている
		fmt.Printf("Fail to read .env file : %v", err)
		return err
	}

	//[Good]  環境変数から接続情報を取得している
	/*[Motto] エラーハンドリングができると良い
	  if dialect := os.Getenv("DIALECT"); dialect == "" {
			return fmt.Errorf("DIALECT value is empty")
		}
	*/
	DBMS := os.Getenv("DIALECT")
	USER := os.Getenv("USER_NAME")
	PASS := os.Getenv("PASSWORD")
	PROTOCOL := os.Getenv("PROTOCOL")
	DBNAME := os.Getenv("DB_NAME")

	/*[Motto] 文字列操作はSprintfを使うと可読性が上がる
	  CONNECT := fmt.Sprintf("%s:%s@%s/%s?parseTime=true", USER, PASS, PROTOCOL, DBNAME)
	*/
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	db, err = gorm.Open(DBMS, CONNECT)

	if err != nil {
		//[Motto] panicはプログラムが終了してしまうため、呼ぶならmain関数で呼ぶ
		return err
	}

	return nil
}
