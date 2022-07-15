package http_errors

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

var (
	ErrUserNotFound = echo.NewHTTPError(http.StatusNotFound, "user with given id does not exist")
	GeneralError    = echo.NewHTTPError(http.StatusBadRequest, "something gone wrong")
)
