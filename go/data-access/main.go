package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	// "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// Capture connection properties.
	// cfg := mysql.Config{
	// 	User:   os.Getenv("DBUSER"),
	// 	Passwd: os.Getenv("DBPASS"),
	// 	Net:    "tcp",
	// 	Addr:   "127.0.0.1:5432",
	// 	DBName: "recordings",
	// }
	// connStr:=cfg.FormatDSN()
	connStr := fmt.Sprintf("user=%s password=secret host=localhost port=5432 database=recordings sslmode=disable", os.Getenv("USER"))

	// Get a database handle.
	var err error
	db, err = sql.Open("pgx", connStr)
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
