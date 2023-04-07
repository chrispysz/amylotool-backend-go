package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB
var server = "amylotool-server.database.windows.net"
var port = 1433
var user = "amylotool"
var password = "Admin123"
var database = "amylotool_core"

func initializeDatabase() {
	defer fmt.Printf("Connected!")

	// Build connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)
	var err error
	// Create connection pool
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
}
