package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/serenade010/gollab/internal/model"
)

var hub map[string]*websocket.Conn

type Image struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Base64   string `json:"source"`
}

func (app *application) createUser(c *gin.Context) {
	var newUser model.User

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	// Insert to DB
	app.users.Insert(newUser.Name)
	c.IndentedJSON(http.StatusCreated, newUser)

}

func (app *application) addFriend(c *gin.Context) {
	var newFriendList model.FriendList

	err := c.BindJSON(&newFriendList)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Insert to DB
	err = app.friendLists.Insert(newFriendList.Me, newFriendList.Friend)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	c.IndentedJSON(http.StatusCreated, newFriendList)

}

func (app *application) getFriendList(c *gin.Context) {
	user := c.Param("user")
	friends, err := app.friendLists.GetList(user)
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, friends)

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
