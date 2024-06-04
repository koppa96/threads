package test

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/lucsky/cuid"
	"go.temporal.io/sdk/client"
	"net/http"
	"threads/temporal/workflows"
)

type StartTestWorkflowResource struct {
	client client.Client
}

func NewStartTestWorkflowResource(client client.Client) *StartTestWorkflowResource {
	return &StartTestWorkflowResource{client: client}
}

func (r *StartTestWorkflowResource) Method() string {
	return http.MethodPost
}

func (r *StartTestWorkflowResource) Path() string {
	return "/api/workflows/test/start"
}

func (r *StartTestWorkflowResource) Handle(c echo.Context) error {
	run, err := r.client.ExecuteWorkflow(
		c.Request().Context(),
		client.StartWorkflowOptions{ID: fmt.Sprintf("test-%s", cuid.New()), TaskQueue: "default"},
		workflows.TestWorkflow,
	)
	if err != nil {
		return fmt.Errorf("failed to start workflow test: %w", err)
	}

	err = run.Get(c.Request().Context(), nil)
	if err != nil {
		return fmt.Errorf("failed to await workflow test: %w", err)
	}

	return c.NoContent(200)
}
