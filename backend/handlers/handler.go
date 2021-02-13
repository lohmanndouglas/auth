package handlers

import (
	"github.com/auth/backend/clock"
	"github.com/auth/backend/providers"
	"github.com/auth/backend/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

type HTTPHandler struct {
	clock     clock.Clock
	providers *providers.Provider
	repository *repository.SqlRepository
}

func NewHTTPHandler() (*HTTPHandler, error) {
	providers, err := providers.NewProvider()
	if err != nil {
		return nil, err
	}
	repo, err := repository.NewSqlRepository()
	if err != nil {
		return nil, err
	}
	httpHandler := &HTTPHandler{
		clock:      &clock.RealClock{},
		providers: providers,
		repository: repo,
	}
	return httpHandler, nil
}

func (h *HTTPHandler) Index(context echo.Context) error {
	return context.JSON(http.StatusOK, "the api is ok")
}