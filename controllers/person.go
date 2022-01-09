package controllers

import (
	"net/http"

	"example/skeleton/structs"
	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetPersons(c *gin.Context) {
	var (
		persons []structs.Person
	)

	idb.db.Find(&persons)

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "",
		"data":    persons,
	})
}

func (idb *InDB) CreatePerson(c *gin.Context) {
	var person structs.Person

	firstName := c.PostForm("first_name")
	lastName := c.PostForm("last_name")

	person.FirstName = firstName
	person.LastName = lastName

	err := idb.db.Create(&person)

	if err.Error != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  false,
			"message": err.Error,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "",
		"data":    nil,
	})
}
