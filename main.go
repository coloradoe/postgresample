package main

import (
	"database/sql"
	"fmt"
	"net/url"
	"strconv"

	_ "github.com/lib/pq"
)

func main() {

	//r := InsertValue()

	//fmt.Println(r)

	UpdateValue()

	DeleteValue()
}

// GetPostgresURL builds the PostgreSQL URL
func GetPostgresURL(username string, password string, host string, dbname string, sslmode bool) string {
	v := url.Values{}

	if sslmode {
		v.Set("sslmode", "require")
	} else {
		v.Set("sslmode", "disable")
	}

	u := &url.URL{
		Scheme:   "postgresql",
		User:     url.UserPassword(username, password),
		Host:     host,
		Path:     "/" + dbname,
		RawQuery: v.Encode(),
	}

	return u.String()
}

// InsertValue writes a value to the database
func InsertValue() string {

	psqlInfo := GetPostgresURL("postgres", "", "localhost:5432", "app_demo", false)

	// Validating arguments provided to psqlInfo
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Verifying connection to the Database
	fmt.Println("Verifying Connection to the database...")
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection Verified!")

	// SQL statement to insert values to DB
	sqlStatement := `
     INSERT INTO users (age, email, first_name, last_name)
     VALUES ($1, $2, $3, $4)
     RETURNING id`
	id := 0
	// By using the function QueryRow() instead of Exec() we are telling the db to return one row of data
	// it being the id of our newly created record
	err = db.QueryRow(sqlStatement, 35, "harley@queen.com", "Harley", "Queen").Scan(&id)
	if err != nil {
		panic(err)
	}

	// Converting the record ID to a string
	a := strconv.Itoa(id)

	return "New record ID is:" + a
}

// UpdateValue updates values in the DB
func UpdateValue() {

	psqlInfo := GetPostgresURL("postgres", "", "localhost:5432", "app_demo", false)

	// Validating arguments provided to psqlInfo
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Verifying connection to the Database
	fmt.Println("Verifying Connection to the database...")
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection Verified!")

	fmt.Println("Updating DB Value")
	sqlStatement := `
      UPDATE users
      SET first_name = $2, last_name = $3
      WHERE id = $1
      RETURNING id, email;`
	var email string
	var id int
	err = db.QueryRow(sqlStatement, 10, "Darth", "Vader").Scan(&id, &email)
	if err != nil {
		panic(err)
	}
	fmt.Println(id, email)
	fmt.Println("Successfully updated DB Value!")

}

// DeleteValue will remove a value from the Database
func DeleteValue() {
	psqlInfo := GetPostgresURL("postgres", "", "localhost:5432", "app_demo", false)

	// Validating arguments provided to psqlInfo
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Verifying connection to the Database
	fmt.Println("Verifying Connection to the database...")
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection Verified!")
	fmt.Println("Deleting Row")
	sqlStatement := `DELETE FROM users WHERE id = $1 RETURNING id, email;`
	var email string
	var id int
	err = db.QueryRow(sqlStatement, 12).Scan(&id, &email)
	if err != nil {
		panic(err)
	}
	fmt.Println(id, email)
	fmt.Println("Row Successfully Deleted!")
}
