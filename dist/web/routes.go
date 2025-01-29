package main

func (app *application) routes() {

	app.router.Use(app.allowedCors())
	app.router.POST("/login", app.loginPost)
	app.router.POST("/signup", app.signUpPost)
	protected := app.router.Group("/")
	protected.Use(app.isAuthorized)
	protected.GET("/", app.home)

}
