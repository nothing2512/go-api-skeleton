package system

import (
	"example/skeleton/controllers"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	c := controllers.NewInstance(controllers.DbConfig{
		Username: "admin",
		Password: "",
		Host:     "localhost",
		Port:     "3306",
		Dbname:   "test",
	})

	router.GET("/persons", c.GetPersons)
	router.POST("/person", c.CreatePerson)
}
