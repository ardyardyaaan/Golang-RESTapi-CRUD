package main

import (
	"./config"
	"./controllers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/employee/:id", inDB.GetEmployee)
	router.GET("/employees", inDB.GetEmployees)
	router.POST("/employee", inDB.CreateEmployee)
	router.PUT("/employee/:id", inDB.UpdateEmployee)
	router.DELETE("/employee/:id", inDB.DeleteEmployee)
	router.Run(":8080")
}
