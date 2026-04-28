package main

import (
	"fmt"
	"fullstack-backend/utils"
)

func main() {
	hash, err := utils.HashPassword("admin123")
	if err != nil {
		panic(err)
	}
	fmt.Println("Password hash:", hash)
}
