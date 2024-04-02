// db/database.go

package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	return sql.Open("mysql", "root:@tcp(localhost:3306)/myTasks")
}
