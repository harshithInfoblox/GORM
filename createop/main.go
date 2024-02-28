package main

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Define Model Struct
type User struct {
	ID        uint
	Name      string
	Age       int
	Birthday  time.Time
	CreatedAt time.Time // Add CreatedAt field
}

func main() {
	// Connect to SQLite Database
	db, err := gorm.Open(sqlite.Open("/Users/hp/go/src/gormtutorial/createop/test.db"), &gorm.Config{CreateBatchSize: 100})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("failed to migrate database")
	}

	// Create a Single Record
	user := User{Name: "1stNormal", Age: 18, Birthday: time.Now()}
	result := db.Create(&user)
	if result.Error != nil {
		panic(result.Error) // Handle error
	}

	// Create Multiple Records
	users := []*User{
		{Name: "multiRec1", Age: 18, Birthday: time.Now()},
		{Name: "multiRec2", Age: 19, Birthday: time.Now()},
	}
	result = db.Create(users)
	if result.Error != nil {
		panic(result.Error) // Handle error
	}

	// Example of Select and Omit
	db.Select("Name", "Age", "CreatedAt").Create(&user)
	// This will execute: INSERT INTO `users` (`name`,`age`) VALUES ("Jinzhu", 18)

	db.Omit("Name", "Age").Create(&user)

	users = []*User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	db.Create(&users)

	// This will execute: INSERT INTO `users` (`birthday`) VALUES ("2020-01-01 00:00:00.000")
	var res User
	db.First(&res, user.ID)
	fmt.Println("Created User:", result)

}
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// Perform actions before creating a new record
	u.Birthday = time.Now() // Set Birthday to current time
	return nil
}
func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	// Perform actions after creating a new record
	fmt.Println("User created successfully!")
	return nil
}
