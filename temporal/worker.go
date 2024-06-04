package temporal

import (
	"context"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.uber.org/fx"
	"threads/temporal/activities"
	"threads/temporal/workflows"
)

func NewWorker(lc fx.Lifecycle, client client.Client, testActivity *activities.TestActivity) worker.Worker {
	w := worker.New(client, "default", worker.Options{})

	w.RegisterWorkflow(workflows.TestWorkflow)
	w.RegisterActivity(testActivity)

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return w.Start()
		},
		OnStop: func(ctx context.Context) error {
			w.Stop()
			return nil
		},
	})

	return w
}
