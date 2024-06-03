package threads

import (
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"threads/ent"
	"threads/ent/message"
	"time"
)

type ThreadDetails struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	LastMessage *time.Time `json:"lastMessage,omitempty"`
	Messages    int        `json:"messages"`
}

type GetByIDResource struct {
	client *ent.Client
}

func NewGetByIDResource(client *ent.Client) *GetByIDResource {
	return &GetByIDResource{client: client}
}

func (r *GetByIDResource) Method() string {
	return http.MethodGet
}

func (r *GetByIDResource) Path() string {
	return "/api/threads/:id"
}

func (r *GetByIDResource) Handle(c echo.Context) error {
	id := c.Param("id")

	thread, err := r.client.Thread.Get(c.Request().Context(), id)
	if err != nil {
		return fmt.Errorf("failed to get thread by id: %v", err)
	}

	lastMessages, err := r.client.Thread.QueryMessages(thread).
		Order(message.ByCreated(sql.OrderDesc())).
		Limit(1).
		All(c.Request().Context())
	if err != nil {
		return fmt.Errorf("failed to get last message: %v", err)
	}

	messageCount, err := r.client.Thread.QueryMessages(thread).
		Count(c.Request().Context())
	if err != nil {
		return fmt.Errorf("failed to get message count: %v", err)
	}

	var lastMessage *time.Time
	if len(lastMessages) > 0 {
		lastMessage = &lastMessages[0].Created
	}

	return c.JSON(http.StatusOK, ThreadDetails{
		ID:          thread.ID,
		Name:        thread.Name,
		LastMessage: lastMessage,
		Messages:    messageCount,
	})
}
