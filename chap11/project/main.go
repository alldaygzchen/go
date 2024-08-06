package main

import (
	"chap11/project/db"
	"chap11/project/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server :=gin.Default()
	// fmt.Println("give me the type of server variable",reflect.TypeOf(server))
	routes.RegisterRoutes(server)
	server.Run(":8080")
}

