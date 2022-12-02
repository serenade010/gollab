package main

import (
	"database/sql"
	"flag"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type application struct {
	router *gin.Engine
	hub    map[string]*websocket.Conn
	// users  *model.UserlModel
}

func main() {

	//settings for application
	app := application{
		router: gin.Default(),
		hub:    make(map[string]*websocket.Conn),
	}
	app.routes()

	err := loadENV()
	if err != nil {
		log.Fatal(err)
	}

	dsn := flag.String("dsn", os.Getenv("DB_URL"), "SQL data source name")
	flag.Parse()

	db, err := openDB(*dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app.router.Run("localhost:8080")
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func loadENV() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	return nil
}
