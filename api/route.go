package api

import (
	"github.com/labstack/echo/v4"
	"github.com/nilotpaul/go-echo-htmx/config"
)

type Routes struct {
	env config.EnvConfig
}

func NewRouter(env config.EnvConfig) *Routes {
	return &Routes{
		env: env,
	}
}

func (r *Routes) RegisterRoutes(router *echo.Router) {
	router.Add("GET", "/", handleGetHome)
}
