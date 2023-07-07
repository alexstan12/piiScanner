package test

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"piiScanner/pkg/pii"

	. "github.com/smartystreets/goconvey/convey"
	_ "github.com/snowflakedb/gosnowflake"
)

func TestColumnNameRegexDetector_Detect(t *testing.T) {
	Convey("Test db connection", t, func() {
		instance := &DBInstance{}

		output := captureOutput(instance.ConnectDB, t)
		So(output, ShouldContainSubstring, "Successfully logged in")
		So(output, ShouldContainSubstring, "ping succeeded!")
		So(instance.db, ShouldNotBeNil)

		defer instance.db.Close()

		Convey("Fetch column names", func() {
			dbName := os.Getenv("DB_NAME")
			So(dbName, ShouldNotEqual, "")

			schemaName := os.Getenv("SCHEMA_NAME")
			So(schemaName, ShouldNotEqual, "")

			piiCount := 0
			columns, err := instance.ColumnNamesQuery(dbName, schemaName)
			So(err, ShouldBeNil)
			detector := &pii.ColumnNameRegexDetector{}
			So(func() {
				for _, column := range columns {
					isPii := detector.Detect(column.Name)
					if isPii {
						t.Logf("PII column name found: %v", column.Name)
						piiCount++
					}
				}
			}, ShouldNotPanic)

			t.Logf("The count of PIIs is %v", piiCount)
		})
	})
}

type DBInstance struct {
	db *sql.DB
}

func (instance *DBInstance) ConnectDB(t *testing.T) {
	credentials := os.Getenv("CREDENTIALS")
	if credentials == "" {
		t.Error("Missing db credentials, can't continue")
	}

	dbType := os.Getenv("DB_TYPE")
	if credentials == "" {
		t.Error("Missing db driver type, can't continue")
	}

	var err error
	instance.db, err = sql.Open(dbType, credentials)
	if err != nil {
		log.Println("Failed to login ", err)
		defer instance.db.Close()
		return
	} else {
		log.Println("Successfully logged in")
	}

	pingErr := instance.db.Ping()
	if pingErr != nil {
		log.Println(pingErr)
	} else {
		log.Println("ping succeeded!")
	}
}

type Column struct {
	Name string
}

// columnNamesQuery queries for names of columns
func (instance *DBInstance) ColumnNamesQuery(dbName, schemaName string) ([]Column, error) {
	// An columns slice to hold data from query
	var columns []Column

	query := fmt.Sprintf("SELECT COLUMN_NAME FROM %v.%v.columns;", dbName, schemaName)
	rows, err := instance.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("columnNamesQuery: %v", err)
	}
	defer rows.Close()
	
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var column Column
		if err := rows.Scan(&column.Name); err != nil {
			return nil, fmt.Errorf("columnNamesQuery: %v", err)
		}
		columns = append(columns, column)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("columnNamesQuery: %v", err)
	}
	return columns, nil
}

func captureOutput(f func(t *testing.T), t *testing.T) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f(t)
	log.SetOutput(os.Stderr)
	return buf.String()
}
