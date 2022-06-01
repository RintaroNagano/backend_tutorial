package model

// import(
// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/mysql"
// )

type User struct {
	UserId int
	Password string
	Token string
}