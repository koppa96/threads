package resources

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type Resource interface {
	Method() string
	Path() string
	Handle(c echo.Context) error
}

func AsResource(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Resource)),
		fx.ResultTags(`group:"routes"`),
	)
}
