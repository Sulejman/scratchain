package controller

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

type Controller struct {
	DB *sql.DB
}
