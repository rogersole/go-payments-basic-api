package apis

import (
	"strconv"

	"github.com/go-ozzo/ozzo-routing"
	"github.com/rogersole/payments-basic-api/util"
)

const (
	DefaultPageSize = 100
	MaxPageSize     = 1000
)

func getPaginatedListFromRequest(c *routing.Context, count int) *util.PaginatedList {
	page := parseInt(c.Query("page"), 1)
	perPage := parseInt(c.Query("per_page"), DefaultPageSize)
	if perPage <= 0 {
		perPage = DefaultPageSize
	}
	if perPage > MaxPageSize {
		perPage = MaxPageSize
	}
	return util.NewPaginatedList(page, perPage, count)
}

func parseInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	if result, err := strconv.Atoi(value); err == nil {
		return result
	}
	return defaultValue
}
