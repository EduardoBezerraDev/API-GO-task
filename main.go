package main

import (
	"database/sql"
	"log"
	"net/http"

	"API-GO-task/v1/api"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/myTasks")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := api.NewRouter(db)

	log.Fatal(http.ListenAndServe(":8080", router))
}
