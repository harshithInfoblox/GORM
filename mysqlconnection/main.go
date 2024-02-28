package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Define a model
type User struct {
	ID   uint
	Name string
	Age  int
}

func main() {
	// MySQL DSN (replace placeholders with your actual MySQL credentials)
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	// Open a new database connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//defer db.Close()
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			panic(err)
		}
		sqlDB.Close()
	}()
	// Auto-migrate the schema
	db.AutoMigrate(&User{})

	// Create a new record
	user := User{Name: "John", Age: 30}
	db.Create(&user)

	// Read records
	var retrievedUser User
	db.First(&retrievedUser, 1) // Get the first user with ID 1

	// Update record
	db.Model(&retrievedUser).Update("Age", 40)

	// Delete record
	db.Delete(&retrievedUser)

	// Now you can use 'db' to interact with the database
}
