package controllers

import (
	"example/skeleton/structs"
	"fmt"
	"github.com/jinzhu/gorm"
	"regexp"
	"strings"
)

type InDB struct {
	db *gorm.DB
}

type DbConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Dbname   string
}

func NewInstance(config DbConfig) InDB {
	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Dbname,
	)

	matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")

	gorm.AddNamingStrategy(&gorm.NamingStrategy{
		DB: func(s string) string {
			snake := matchFirstCap.ReplaceAllString(s, "${1}_${2}")
			snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
			return strings.ToLower(snake)
		},
		Table: func(s string) string {
			snake := matchFirstCap.ReplaceAllString(s, "${1}_${2}")
			snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
			return strings.ToLower(snake)
		},
		Column: func(s string) string {
			snake := matchFirstCap.ReplaceAllString(s, "${1}_${2}")
			snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
			return strings.ToLower(snake)
		},
	})

	db, err := gorm.Open("mysql", dbUri)
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(
		structs.Person{},
	)

	return InDB{
		db: db,
	}
}
