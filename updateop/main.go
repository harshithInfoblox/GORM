package main

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type User struct {
	ID        uint
	Name      string
	Age       uint
	Birthday  time.Time
	UpdatedAt time.Time
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})

	// Create
	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	db.Create(&user)
	fmt.Println("User created successfully!")

	// Update Save All Fields
	db.First(&user)
	user.Name = "Jinzhu 2"
	user.Age = 100
	db.Save(&user)
	fmt.Println("Save All Fields:")
	fmt.Println("Updated User:", user)

	// Update single column
	db.Model(&user).Update("Name", "hello")
	fmt.Println("Update single column:")
	fmt.Println("Updated User:", user)

	// Updates multiple columns
	db.Model(&user).Updates(User{Name: "hello", Age: 18})
	fmt.Println("Updates multiple columns:")
	fmt.Println("Updated User:", user)

	// Update Selected Fields
	db.Model(&user).Select("Name").Updates(map[string]interface{}{"Name": "hello", "Age": 18})
	fmt.Println("Update Selected Fields:")
	fmt.Println("Updated User:", user)

	// Update Hooks
	// Define a BeforeUpdate hook
	// db.Callback().Update().Before("gorm:update").Register("check_role", func(db *gorm.DB) {
	// 	var user User
	// 	db.Find(&user)
	// 	if user.Role == "admin" {
	// 		db.AddError(errors.New("admin user not allowed to update"))
	// 	}
	// })

	// Batch Updates
	db.Model(&User{}).Where("age > ?", 10).Update("Name", "New Name")
	fmt.Println("Batch Updates done!")

	// Block Global Updates
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&User{}).Update("Name", "Global Update")
	fmt.Println("Global Update done!")

	// Returning Data From Modified Rows
	var users []User
	result := db.Model(&users).Clauses(clause.Returning{}).Where("role = ?", "admin").Update("Salary", gorm.Expr("salary * ?", 2))
	fmt.Println("Returning Data From Modified Rows:")
	fmt.Println("Updated records count:", result.RowsAffected)

	// Check Field has changed?
	if db.Model(&user).Where("id = ?", user.ID).Updates(User{Name: "hello"}).RowsAffected == 1 {
		fmt.Println("Name field has changed!")
	}
}
