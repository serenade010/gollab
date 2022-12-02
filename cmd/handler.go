package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var hub map[string]*websocket.Conn

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Image struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Base64   string `json:"source"`
}

var users = []User{
	{ID: "1", Name: "Blue Train"},
	{ID: "2", Name: "Jeru"},
	{ID: "3", Name: "Sarah Vaughan and Clifford Brown"},
}

func (app *application) getUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func (app *application) postUser(c *gin.Context) {
	var newUser User

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	// Add the new album to the slice.
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func (app *application) getUserByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range users {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func (app *application) SocketHandler(c *gin.Context) {
	user := c.Param("user")
	upGrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	// if err != nil {
	// 	panic(err)
	// }
	app.hub[user] = ws

	fmt.Println(app.hub)
	welcomeMsg := fmt.Sprintf("Hello %s", user)
	ws.WriteMessage(1, []byte(welcomeMsg))

	defer func() {
		closeSocketErr := ws.Close()
		delete(hub, user)
		if closeSocketErr != nil {
			panic(err)
		}
	}()

	for {
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			return
		}

		fmt.Printf("Message Type: %d, Message: %s\n", msgType, string(msg))
		// err = ws.WriteJSON(struct {
		// 	Reply string `json:"reply"`
		// }{
		// 	Reply: "you received a message!",
		// })
		// if err != nil {
		// 	panic(err)
		// }
	}
}

func (app *application) sendImage(c *gin.Context) {

	var newImage Image

	if err := c.BindJSON(&newImage); err != nil {
		return
	}

	ws := app.hub[newImage.Receiver]
	ws.WriteMessage(1, []byte("ssss"))
	c.String(http.StatusOK, "image sent")

}
