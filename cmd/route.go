package main

func (app *application) routes() {

	//set up the routes
	app.router.GET("/socket/:user", app.SocketHandler)
	app.router.GET("/user", app.getUser)
	app.router.GET("/user/:id", app.getUserByID)
	app.router.POST("/user", app.postUser)
	app.router.POST("/image", app.sendImage)

}
