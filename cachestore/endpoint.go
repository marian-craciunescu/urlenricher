package cachestore

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/marian-craciunescu/urlenricher/models"
	"net/http"
	"time"
)

type Endpoint interface {
	Save(url *models.URL) error
	Get(baseURL string) (*models.URL, error)
	Delete(baseURL string) error
	Resolve(ctx echo.Context) error
	Dump(ctx echo.Context) error
}

type urlEndpoint struct {
	store *URLCacheStore
}

func NewURLEndpoint(store *URLCacheStore) Endpoint {
	return &urlEndpoint{store}
}

func (e *urlEndpoint) Save(url *models.URL) error {
	return e.store.save(url.Address, url)
}
func (e *urlEndpoint) Get(baseURL string) (*models.URL, error) {
	return e.store.get(baseURL)
}
func (e *urlEndpoint) Delete(baseURL string) error {
	return models.ErrNotImplemented
}

func (e *urlEndpoint) Resolve(ctx echo.Context) error {
	target := ctx.QueryParam("target")
	logger.WithField("target_url", target).Info("Endpoint Resolve")

	url, err := e.store.get(target)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, url)

}

func (e *urlEndpoint) Dump(ctx echo.Context) error {
	d, err := e.store.Dump()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"dump_ts":         time.Now().UTC().Format(time.RFC3339),
		"total_dump_size": fmt.Sprintf("%d", d),
	})
}
