package main

import (
	"github.com/labstack/echo/v4"
	"go.temporal.io/sdk/worker"
	"go.uber.org/fx"
	"threads/deps"
	"threads/resources"
	"threads/resources/threads"
	"threads/resources/workflows/test"
	"threads/temporal"
	"threads/temporal/activities"
)

func main() {
	fx.New(
		fx.Provide(
			deps.NewConfig,
			deps.NewClient,
			activities.NewTestActivity,
			temporal.NewClient,
			temporal.NewWorker,
			resources.AsResource(threads.NewGetAllResource),
			resources.AsResource(threads.NewGetByIDResource),
			resources.AsResource(threads.NewCreateThreadResource),
			resources.AsResource(test.NewStartTestWorkflowResource),
			fx.Annotate(
				deps.NewEcho,
				fx.ParamTags("", "", `group:"routes"`),
			),
		),
		fx.Invoke(func(e *echo.Echo, w worker.Worker) {}),
	).Run()
}
