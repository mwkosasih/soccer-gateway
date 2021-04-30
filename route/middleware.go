package route

import (
	"github.com/labstack/echo/v4"
	"github.com/mwkosasih/soccer-gateway/route/middleware"
)

var middlewareHandler = map[string]echo.MiddlewareFunc{
	"api_key": middleware.AuthCheckAppKey,
}
