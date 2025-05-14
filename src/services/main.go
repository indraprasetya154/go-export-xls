package main

import (
	"github.com/indraprasetya154/go-export-xls/src/app"
	"github.com/indraprasetya154/go-export-xls/src/router" // Import the router package
	"github.com/spf13/viper"
)

func main() {
	app := app.InitApp()
	// Create a new router instance
	e := router.NewRouter(app.Echo)

	// Start server
	port := viper.GetString("APP_PORT")
	e.Start(":" + port)
}
