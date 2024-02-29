// package main

// import (
// 	"log"
// 	"time"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// type User struct {
// 	ID        uint
// 	Name      string
// 	Age       int
// 	Birthday  time.Time
// 	CreatedAt time.Time // Add CreatedAt field
// }

// func main() {
// 	// Define PostgreSQL DSN
// 	dsn := "host=localhost user='postgress' password='seceret' dbname=mydatabase port=5432 sslmode=disable TimeZone=Asia/Shanghai"

// 	// Initialize GORM database instance with pgx driver and configuration
// 	config := postgres.Config{
// 		DSN:                  dsn,
// 		PreferSimpleProtocol: true, // Disables implicit prepared statement usage
// 	}

// 	db, err := gorm.Open(postgres.New(config), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Error connecting to database: %v", err)
// 	}

// 	err = db.AutoMigrate(&User{})
// 	if err != nil {
// 		panic("failed to migrate database")
// 	}

//		// Create a Single Record
//		user := User{Name: "1stNormal", Age: 18, Birthday: time.Now()}
//		result := db.Create(&user)
//		if result.Error != nil {
//			panic(result.Error) // Handle error
//		}
//		//defer db.Close() // Remember to close the database connection when finished
//	}
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Company struct {
	comp_ID int
	name    string
	city    string
	domain  string
}
type Employees struct {
	emp_ID  int
	name    string
	dept_id int
	proj_id int
}
type Project struct {
	proj_id   int
	proj_name string
}

func (c *Company) CreateTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS COMPANY(
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        city VARCHAR(100) NOT NULL,
        domain VARCHAR(100) NOT NULL
    )`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
func (c *Company) InsertValues(db *sql.DB) {
	query := `INSERT INTO COMPANY(name, city, domain) VALUES($1, $2, $3) RETURNING id`
	err := db.QueryRow(query, c.name, c.city, c.domain).Scan(&c.comp_ID)
	if err != nil {
		log.Fatal(err)
	}
}
func (c *Company) GetData(db *sql.DB) {
	rows, err := db.Query("SELECT id, name, city, domain FROM COMPANY")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&c.comp_ID, &c.name, &c.city, &c.domain)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Company ID : %d\n Company Name : %s\n Company Location : %s\n Company Domain : %s\n", c.comp_ID, c.name, c.city, c.domain)
	}
	fmt.Print("\n\n")
}
func (e *Employees) CreateTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS EMPLOYEES(
        emp_id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        dept_id INTEGER NOT NULL,
        proj_id INTEGER NOT NULL
    )`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
func (e *Employees) InsertValues(db *sql.DB) {
	query := `INSERT INTO EMPLOYEES(name, dept_id, proj_id) VALUES($1, $2, $3) RETURNING emp_id`
	err := db.QueryRow(query, e.name, e.dept_id, e.proj_id).Scan(&e.emp_ID)
	if err != nil {
		log.Fatalf("Error inserting values into EMPLOYEES table: %v", err)
	}
	log.Println("Employee values inserted successfully")
}
func (e *Employees) GetData(db *sql.DB) {
	rows, err := db.Query("SELECT emp_id, name, dept_id, proj_id FROM EMPLOYEES")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&e.emp_ID, &e.name, &e.dept_id, &e.proj_id)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Employee ID : %d\n Employee Name : %s\n Department ID : %d\n Project ID : %d", e.emp_ID, e.name, e.dept_id, e.proj_id)
	}
	fmt.Print("\n\n")
}
func (p *Project) CreateTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS PROJECT(
        proj_id SERIAL PRIMARY KEY,
        proj_name VARCHAR(100) NOT NULL
    )`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
func (p *Project) InsertValues(db *sql.DB) {
	query := `INSERT INTO PROJECT(proj_name) VALUES($1) RETURNING proj_id`
	err := db.QueryRow(query, p.proj_name).Scan(&p.proj_id)
	if err != nil {
		log.Fatalf("Error inserting values into PROJECT table: %v", err)
	}
	log.Println("Project values inserted successfully")
}
func (p *Project) GetData(db *sql.DB) {
	rows, err := db.Query("SELECT proj_id, proj_name FROM PROJECT")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&p.proj_id, &p.proj_name)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Project ID : %d\n Project Name : %s", p.proj_id, p.proj_name)
	}
	fmt.Print("\n\n")
}

// Interface
type organization interface {
	CreateTable(db *sql.DB)
	InsertValues(db *sql.DB)
	GetData(db *sql.DB)
}

func crud(o organization, db *sql.DB) {
	o.CreateTable(db)
	o.InsertValues(db)
	o.GetData(db)
}
func main() {
	connstr := "postgres://postgres:secret@localhost:5432/mydatabase?sslmode=disable"
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Create an instance of company table
	company := &Company{name: "Amazon", city: "Delhi", domain: "Retail and Cloud"}
	crud(company, db) // calling crud function
	// Create an instance of Employees table
	employees := &Employees{name: "Sanju", dept_id: 6, proj_id: 6}
	crud(employees, db) // calling crud function
	// Create an instance of Project table
	project := &Project{proj_name: "Order Checkout Automation"}
	crud(project, db) // calling crud function
}
