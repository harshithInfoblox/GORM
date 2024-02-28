package main

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Define Model Struct
type User struct {
	ID       uint
	Name     string
	Age      int
	Birthday time.Time
}

func main() {
	// Connect to SQLite Database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto Migrate the User model
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("failed to migrate database")
	}

	// Create a Single Record
	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	result := db.Create(&user)
	if result.Error != nil {
		panic("failed to create user")
	}
	fmt.Println("User created successfully!")

	// Retrieve a Single Object
	var retrievedUser User
	db.First(&retrievedUser)
	fmt.Println("Retrieved User:", retrievedUser)

	// Retrieve Objects with Primary Key
	var users []User
	db.Find(&users, []uint{user.ID})
	fmt.Println("Retrieved Users:", users)

	// Conditions
	var filteredUser User
	db.Where("name = ?", "Jinzhu").First(&filteredUser)
	fmt.Println("Filtered User:", filteredUser)

	// Selecting Specific Fields
	var selectedUser User
	db.Select("name", "age").Find(&selectedUser)
	fmt.Println("Selected User:", selectedUser)

	// Order, Limit & Offset
	var orderedUsers []User
	db.Order("age desc").Limit(2).Offset(1).Find(&orderedUsers)
	fmt.Println("Ordered Users:", orderedUsers)

	// Group By & Having
	var groupedResult struct {
		Name  string
		Total int
	}
	db.Model(&User{}).Select("name, sum(age) as total").Group("name").First(&groupedResult)
	fmt.Println("Grouped Result:", groupedResult)

	// Joins
	var joinResult struct {
		Name  string
		Email string
	}
	db.Model(&User{}).Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&joinResult)
	fmt.Println("Join Result:", joinResult)

	// Scan
	var scanResult struct {
		Name string
		Age  int
	}
	db.Table("users").Select("name", "age").Where("name = ?", "Jinzhu").Scan(&scanResult)
	fmt.Println("Scan Result:", scanResult)
}
