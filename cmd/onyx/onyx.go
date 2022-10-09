package main

import (
	"flag"
	"net/http"

	"github.com/ShapleyIO/onyx/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func setup() error {
	cfg := config.LoadDefaultConfig()

	flag.StringVar(&cfg.SomeFlag, "someflag", "help message", "")

	return nil
}

func onyxRoot(ctx echo.Context) error {
	return ctx.JSONBlob(http.StatusOK, []byte(`{"msg": "Hello, World!"}`))
}

func main() {
	echoServer := echo.New()

	echoServer.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	echoServer.GET("/", onyxRoot)

	echoServer.Logger.Fatal(echoServer.Start(":8080"))
}
