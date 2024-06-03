package threads

import (
	"entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"go.uber.org/fx"
	"net/http"
	"threads/ent"
	"threads/ent/message"
	"threads/resources"
	"time"
)

type ThreadItem struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	LastMessage *time.Time `json:"lastMessage"`
}

type GetAllResource struct {
	client *ent.Client
}

func NewGetAllResource(_ fx.Lifecycle, client *ent.Client) *GetAllResource {
	return &GetAllResource{client: client}
}

func (r *GetAllResource) Method() string {
	return http.MethodGet
}

func (r *GetAllResource) Path() string {
	return "/api/threads"
}

func (r *GetAllResource) Handle(c echo.Context) error {
	pagination, err := resources.ParsePagination(c)
	if err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	count, err := r.client.Thread.Query().
		Count(c.Request().Context())
	if err != nil {
		return fmt.Errorf("failed to count threads : %w", err)
	}

	entities, err := r.client.Thread.Query().
		Offset((pagination.Page - 1) * pagination.PageSize).
		Limit(pagination.PageSize).
		WithMessages(func(query *ent.MessageQuery) {
			query.Order(message.ByID(sql.OrderDesc())).
				Limit(1)
		}).
		All(c.Request().Context())
	if err != nil {
		return fmt.Errorf("failed to get threads list: %w", err)
	}

	return c.JSON(http.StatusOK, resources.PagedList[ThreadItem]{
		Page:     pagination.Page,
		PageSize: pagination.PageSize,
		Total:    count,
		Items: lo.Map(entities, func(item *ent.Thread, index int) ThreadItem {
			var lastMessage *time.Time
			if len(item.Edges.Messages) > 0 {
				lastMessage = &item.Edges.Messages[0].Created
			}

			return ThreadItem{
				ID:          item.ID,
				Name:        item.Name,
				LastMessage: lastMessage,
			}
		}),
	})
}
