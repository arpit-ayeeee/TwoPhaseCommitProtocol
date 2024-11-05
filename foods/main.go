package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var dsn string

func init() {
	dbUser := "root"
	dbPass := "Root@123"
	dbHost := "127.0.0.1"
	dbName := "zomato2pc"
	// Connection string to the DB
	dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)
}

func main() {
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("error pinging the DB.")
	}
	defer DB.Close()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/food/reserve", reserveFoodHandler)

	r.POST("/food/book", bookFoodHandler)

	// start server
	fmt.Println("Food reservation server running on port 8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatal("Unable to start:", err)
	}
}
