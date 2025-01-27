package main

func (app *application) routes() {

	app.router.Use(app.allowedCors())

}
