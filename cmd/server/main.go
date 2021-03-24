package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/abkucuk/go-rest-api/internal/transport/http"
)

// App - the struct which contains things like pointers
// to DB connections
type App struct{}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting up our app")
	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()
	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("failed up to set up server")
		return err
	}
	return nil
}
func main() {
	fmt.Println("Go rest api")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
