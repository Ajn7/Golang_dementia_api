package main

import (
	//"errors"

	"dementia/go_crud/functions"
	"dementia/go_crud/initilizers"
	"github.com/gin-gonic/gin"
)
func init(){
	initilizers.LoadEnvVariable()
	initilizers.ConnectionToDB()
}

func main() {

	router := gin.Default()
	router.POST("/user", functions.UserCreate)//create
	router.POST("/showuser", functions.ShowUser)//showuser
	router.GET("/user/:email/:password", functions.GetUser)
	// router.GET("/user/:id", functions.UserShow)
	router.Run()
}