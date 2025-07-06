package main

import (
	"fmt"
	"chatting-app-mini-version/database"


	 "github.com/gin-gonic/gin"
)

func main(){

	database.ConnectDB()
	r := *gin.Default()
	







		r.Run(":8083")

}