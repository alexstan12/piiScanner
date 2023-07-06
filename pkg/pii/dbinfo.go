package pii

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func connectDB() {
	var err error
	db, err = sql.Open("sqlite3", "path/to/your/database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

func columnNamesByTable() {
	tableName := "your_table_name"
	query := "SELECT COLUMN_NAME FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = ?"

	rows, err := db.Query(query, tableName)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	columnNames := make([]string, 0)

	for rows.Next() {
		var columnName string
		err := rows.Scan(&columnName)
		if err != nil {
			log.Fatal(err)
		}
		columnNames = append(columnNames, columnName)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Column names:")
	for _, columnName := range columnNames {
		fmt.Println(columnName)
	}
}
