package main

import (
	"fmt"
	"fullstack-backend/config"
	"fullstack-backend/database"
	"fullstack-backend/model"
	"fullstack-backend/utils"
)

func main() {
	if err := database.InitDB(); err != nil {
		panic(err)
	}

	hash, err := utils.HashPassword("admin123")
	if err != nil {
		panic(err)
	}
	fmt.Println("Password hash:", hash)

	admin := model.User{
		Username: "admin",
		Password: string(hash),
		Role:     "admin",
	}

	if err := config.DB.Create(&admin).Error; err != nil {
		fmt.Println("Insert failed:", err)
	} else {
		fmt.Println("Admin user created successfully! ID:", admin.ID)
	}
}