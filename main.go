package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func startServer() {
	//Connect the DB
	db, err := sql.Open("mysql", "root:DroneBreaker55@tcp(localhost:3306)/ledgerly")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	e := echo.New()

	e.GET("/api/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "pong from the Ledgerly server",
		})
	})

	// Run server in goroutine so it doesn't block the app
	go func() {
		fmt.Println("✅ Echo API listening on http://localhost:1323")

		if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed {
			fmt.Println("Echo error:", err)
		}
	}()
}

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Start the echo server
	startServer()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Ledgerly",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		// BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
