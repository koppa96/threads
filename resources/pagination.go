package resources

import "github.com/labstack/echo/v4"

type Pagination struct {
	Page     int `query:"page"`
	PageSize int `query:"pageSize"`
}

type PagedList[T any] struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
	Items    []T `json:"items"`
}

func ParsePagination(c echo.Context) (Pagination, error) {
	var pagination Pagination
	if err := c.Bind(&pagination); err != nil {
		return Pagination{}, err
	}

	if pagination.Page < 1 {
		pagination.Page = 1
	}

	if pagination.PageSize < 1 {
		pagination.PageSize = 10
	}

	return pagination, nil
}
