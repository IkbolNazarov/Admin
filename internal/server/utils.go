package server

import (
	"admin/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GeneratePaginationFromRequest(c *gin.Context) models.Pagination {
	limit := 10
	page := 1
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		}
	}
	return models.Pagination{
		Limit: limit,
		Page:  page,
	}
}

func (r *Handler) TotalPageUserInfo(limit int64, length int64) (int64, error) {
	if limit == 0 {
		limit = 10
	}

	totalPage := length / limit
	if length%limit != 0 {
		totalPage++
	}
	return totalPage, nil
}
