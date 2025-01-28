package main

func (app *application) routes() {

	app.router.Use(app.allowedCors())
	app.router.GET("/", app.home)
	app.router.POST("/signUp", app.signUpPost)

}
