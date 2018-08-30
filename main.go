package main

import (
	"github.com/snowspice/go-gin-gorm-router/api/router"
	orm "github.com/snowspice/go-gin-gorm-router/api/database"
)

func main() {
	defer orm.Eloquent.Close()
	router := router.InitRouter()
	router.Run(":8000")
}