package rest

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/marian-craciunescu/urlenricher/cachestore"
	conf "github.com/marian-craciunescu/urlenricher/config"
	"github.com/sirupsen/logrus"
	"net/http"
)

type API interface {
	Start() error
	Stop() error
}

//API is the rest api server struct
type APIServer struct {
	log    *logrus.Entry
	config *conf.Config
	echo   *echo.Echo
	stopCh chan bool
	cache  cachestore.Endpoint
}

func NewAPIServer(config *conf.Config, cacheEndpoint cachestore.Endpoint) API {
	// create the api
	api := APIServer{
		config: config,
		log:    logger,
	}

	// add the endpoints
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/info", api.info)
	e.GET("/resolve", api.cache.Resolve)

	api.echo = e
	return &api
}

func (api *APIServer) info(ctx echo.Context) error {
	logger.WithField("client_ip", ctx.Request().Host).Debug("GET /info")

	return ctx.JSON(http.StatusOK, map[string]string{
		"description": "Urlenricher info page",
		"info":        "demo",
	})
}

func (api *APIServer) Start() (err error) {
	go func() {
		err = api.echo.Start(fmt.Sprintf(":%d", api.config.ServerPort))
		if err != nil {
			logger.WithError(err).Error("Error starting rest api server")
		}
	}()

	return
}

// Stop will shutdown the engine internally
func (api *APIServer) Stop() error {
	logger.Info("Stopping rest api server")
	return api.echo.Close()
}
