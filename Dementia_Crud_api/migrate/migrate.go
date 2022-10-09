package main

import (
	"dementia/go_crud/initilizers"
	"dementia/go_crud/models"
)

func init() {
	initilizers.LoadEnvVariable()
	initilizers.ConnectionToDB()
}
func main() {
	initilizers.DB.AutoMigrate(&models.User{})
}
