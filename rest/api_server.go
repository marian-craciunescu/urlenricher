package rest

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	conf "github.com/marian-craciunescu/urlenricher/config"
	"github.com/sirupsen/logrus"
	"net/http"
)

type API struct {
	log    *logrus.Entry
	config *conf.Config
	echo   *echo.Echo
}

func NewAPIServer(config *conf.Config) *API {
	// create the api
	api := &API{
		config: config,
		log:    logger,
	}

	// add the endpoints
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/info", api.Info)

	api.echo = e
	return api
}

func (api *API) Info(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]string{
		"description": "Urlenricher info page",
		"info":        "demo",
	})
}

func (api *API) Start() error {
	return api.echo.Start(fmt.Sprintf(":%d", api.config.ServerPort))
}

// Stop will shutdown the engine internally
func (api *API) Stop() error {
	return api.echo.Close()
}
