package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"threads/resources"
)

func NewEcho(lc fx.Lifecycle, config *Config, resources []resources.Resource) *echo.Echo {
	e := echo.New()

	for _, resource := range resources {
		e.Add(resource.Method(), resource.Path(), resource.Handle)
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.BindPort)))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})

	return e
}
