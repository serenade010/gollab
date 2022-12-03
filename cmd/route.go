package main

func (app *application) routes() {

	//set up the routes
	app.router.GET("/socket/:user", app.SocketHandler)
	app.router.POST("/user", app.createUser)
	app.router.POST("/friend", app.addFriend)
	app.router.POST("/image", app.sendImage)
	app.router.GET("/friendlist/:user", app.getFriendList)

}
