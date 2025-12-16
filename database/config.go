package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const DB_FILE = "/database/data/dip_alert.db"

func OpenDB() *sql.DB {
    db, err := sql.Open("sqlite3", DB_FILE)
    if err != nil {
        log.Fatalf("Error opening database: %v", err)
    }

    return db
}

func InitializeDB(db *sql.DB)  {
    sqlStmt, err := os.ReadFile("/database/schema.sql")
    if err != nil {
        log.Fatalf("Error reading schema file: %v", err)
    }

    _, err = db.Exec(string(sqlStmt))
    if err != nil {
        log.Printf("%q: %s\n", err, sqlStmt)
    }
}