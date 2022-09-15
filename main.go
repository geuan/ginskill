package main

import (
	"ginskill/common"
	"ginskill/dbs"
	"ginskill/handlers"
	_ "ginskill/validators"
	"github.com/gin-gonic/gin"
)

func main() {
	dbs.InitDB()
	dbs.InitTable()
	r := gin.New()
	r.Use(common.ErrorHandler())
	//r.POST("/users", handlers.UserList)
	r.GET("/users", handlers.UserList)
	r.GET("/users/:id", handlers.UserDetail)
	r.POST("/users", handlers.UserSave)
	r.Run(":8080")
}
