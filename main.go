package main

import (
	"github.com/gin-gonic/gin"

	"sample/controller"
	"sample/model"
	"sample/db"
)



func main() {
	db.GormConnect()
	defer db.DB.Close()

	// Migrate the schema
	db.DB.AutoMigrate(&model.User{})

	r := gin.Default()
	r.GET("/ping", controller.PingHandler)
	r.POST("/signin", controller.SigninHandler)
	r.POST("/signup", controller.SignupHandler)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
