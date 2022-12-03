package main

import (
	"flag"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
	"github.com/serenade010/gollab/internal/model"
)

type application struct {
	router      *gin.Engine
	hub         map[string]*websocket.Conn
	users       *model.UserlModel
	friendLists *model.FriendListlModel
}

func main() {

	//Loading envs
	err := loadENV()
	if err != nil {
		log.Fatal(err)
	}
	dsn := flag.String("dsn", os.Getenv("DB_URL"), "SQL data source name")
	flag.Parse()

	//settings for DB
	db, err := openDB(*dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	//settings for application
	app := application{
		router:      gin.Default(),
		hub:         make(map[string]*websocket.Conn),
		users:       &model.UserlModel{DB: db},
		friendLists: &model.FriendListlModel{DB: db},
	}
	app.routes()
	app.router.Run("localhost:8080")

}
