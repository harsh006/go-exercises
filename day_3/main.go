package main

import (
	"awesomeProject/go-exercises/day_3/Config"
	"awesomeProject/go-exercises/day_3/Models"
	"awesomeProject/go-exercises/day_3/Routes"
	"fmt"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})

	r := Routes.SetupRouter()
	//running
	r.Run()
}
