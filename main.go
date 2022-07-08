package main

import (
	"github.com/labstack/echo/v4"
	"github.com/phuslu/log"
	"github.com/rostikts/social_network/config"
	"github.com/rostikts/social_network/infrastructure/datastore"
	"net/http"
)

func main() {
	e := echo.New()
	cfg := config.NewConfig()
	db := datastore.NewDB(cfg.Database)
	log.DefaultLogger.Debug().Msg("db is initialized")
	defer db.Close()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":3333"))
}
