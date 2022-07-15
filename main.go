package main

import (
	"github.com/labstack/echo/v4"
	"github.com/phuslu/log"
	"github.com/rostikts/social_network/config"
	"github.com/rostikts/social_network/domain/model"
	"github.com/rostikts/social_network/infrastructure/datastore"
	"github.com/rostikts/social_network/pkg/user/repository"
	service2 "github.com/rostikts/social_network/pkg/user/service"
	"net/http"
)

func main() {
	e := echo.New()
	cfg := config.NewConfig()
	db := datastore.NewDB(cfg.Database)
	log.DefaultLogger.Debug().Msg("db is initialized")
	defer db.Close()

	repo := repository.NewUserRepository(db)
	service := service2.NewUserService(repo)
	usr := model.User{
		ID:        1,
		UserName:  "2",
		FirstName: "45555",
		LastName:  "66666",
		Email:     "2512",
		Password:  "521",
	}

	e.GET("/", func(c echo.Context) error {
		_, err := service.UpdateUserData(usr)
		return c.JSON(http.StatusOK, err)
	})
	e.Logger.Fatal(e.Start(":3333"))
}
