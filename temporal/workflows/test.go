package workflows

import (
	"go.temporal.io/sdk/workflow"
	"threads/temporal/activities"
	"time"
)

func TestWorkflow(ctx workflow.Context) error {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		ScheduleToCloseTimeout: 60 * time.Second,
	})

	var a *activities.TestActivity
	return workflow.ExecuteActivity(ctx, a.Test).Get(ctx, nil)
}
