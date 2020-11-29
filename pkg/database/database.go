package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

var (
	// DB is shared database
	DB  *sql.DB
	err error
	// DBErr is database connection error
	DBErr error
)

// Setup opens a database and saves the reference to `Database` struct.
func Setup() {
	var db = DB

	//config := config.GetConfig()

	db, err = sql.Open("sqlite3", "scratchain.db")

	if err != nil {
		DBErr = err
		fmt.Println("db err: ", err)
	}

	DB = db
}

// Genesis creates tables
func Genesis(db *sql.DB) {
	createBlocksTableSQL := `CREATE TABLE blocks (
		"header" TEXT NOT NULL PRIMARY KEY,		
		"previous" TEXT,
		"next" TEXT,
		"transactions" TEXT 		
	  );`

	createTransactionsTableSQL := `CREATE TABLE transactions (
		"from"  TEXT NOT NULL PRIMARY KEY,		
		"to" TEXT,
		"amount" DOUBLE
	  );`

	log.Println("Creating blocks table...")
	statement, err := db.Prepare(createBlocksTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("Blocks table created")

	log.Println("Creating transactions table...")
	statement, err = db.Prepare(createTransactionsTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("Transactions table created")
}

// GetDB helps you to get a connection
func GetDB() *sql.DB {
	return DB
}

// GetDBErr helps you to get a connection
func GetDBErr() error {
	return DBErr
}
