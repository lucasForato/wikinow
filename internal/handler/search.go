package handler

import (
	"net/http"
	"wikinow/component"
	"wikinow/internal/utils"

	"github.com/labstack/echo/v4"
)

func GETSearch(c echo.Context) error {
  return utils.Render(c, http.StatusOK ,component.SearchModal())
}
