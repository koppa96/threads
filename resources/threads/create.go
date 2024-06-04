package threads

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"threads/ent"
)

type CreateThread struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateThreadResource struct {
	client *ent.Client
}

func NewCreateThreadResource(client *ent.Client) *CreateThreadResource {
	return &CreateThreadResource{client: client}
}

func (r *CreateThreadResource) Method() string {
	return http.MethodPost
}

func (r *CreateThreadResource) Path() string {
	return "/api/threads"
}

func (r *CreateThreadResource) Handle(c echo.Context) error {
	var input CreateThread
	if err := c.Bind(&input); err != nil {
		return err
	}

	entity, err := r.client.Thread.Create().
		SetName(input.Name).
		SetDescription(input.Description).
		Save(c.Request().Context())
	if err != nil {
		return fmt.Errorf("failed to create thread: %w", err)
	}

	c.Response().Header().Set("Location", fmt.Sprintf("/api/threads/%s", entity.ID))
	return c.JSON(http.StatusCreated, ThreadDetails{
		ID:   entity.ID,
		Name: entity.Name,
	})
}
