package main

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"threads/resources"
	"threads/resources/threads"
)

func main() {
	fx.New(
		fx.Provide(
			NewConfig,
			NewClient,
			resources.AsResource(threads.NewGetAllResource),
			resources.AsResource(threads.NewGetByIDResource),
			fx.Annotate(
				NewEcho,
				fx.ParamTags("", "", `group:"routes"`),
			),
		),
		fx.Invoke(func(e *echo.Echo) {}),
	).Run()
}
