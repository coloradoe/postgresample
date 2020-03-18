package main

import (
	"database/sql"
	"fmt"
	"net/url"
	"strconv"

	_ "github.com/lib/pq"
)

func main() {

	r := InsertValueToDB()

	fmt.Println(r)
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

// InsertValueToDB writes a value to the database
func InsertValueToDB() string {

    psqlInfo := GetPostgresURL(<role>, <password>, <host:port>, <db_name>, <ssl_mode>)

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

	sqlStatement := `
     INSERT INTO users (age, email, first_name, last_name)
     VALUES ($1, $2, $3, $4)
     RETURNING id`
	id := 0
	// By using the function QueryRow() instead of Exec() we are telling the db to return one row of data
	// it being the id of our newly created record
	err = db.QueryRow(sqlStatement, 75, "email@example.com", "Email", "Example").Scan(&id)
	if err != nil {
		panic(err)
	}

	a := strconv.Itoa(id)

	return "New record ID is:" + a
}
