package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID        uint
	Name      string
	DeletedAt gorm.DeletedAt // Soft delete field
}
type runner struct {
	ID        uint
	Name      string
	Deletedat gorm.DeletedAt
}

func main() {
	// Connect to SQLite database
	db, err := gorm.Open(sqlite.Open("/Users/hp/go/src/gormtutorial/deleteop/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto Migrate the User schema
	db.AutoMigrate(&User{})
	db.AutoMigrate(&runner{})
	// Create a new user
	db.Delete(&runner{}, 1)
	db.Delete(&runner{}, 2)

	db.Unscoped().Where("Deletedat IS NOT NULL").Delete(&runner{})

	db.Create(&runner{ID: 1, Name: "John"})
	db.Create(&runner{ID: 2, Name: "harshith"})

	// Soft delete the user
	db.Delete(&runner{}, 1)

	// // Query soft deleted records
	var users []runner
	db.Unscoped().Find(&users)

	for _, user := range users {
		fmt.Println(user)
	}

	// fmt.Print("**************************\n ")
	db.Unscoped().Where("Deletedat IS NOT NULL").Delete(&runner{})
	var runners []runner
	db.Unscoped().Find(&runners)

	for _, user := range runners {
		fmt.Println(user)
	}
}
