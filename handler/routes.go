package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(v1 *echo.Group) {
	repos := v1.Group("/repos")
	repos.POST("/metrics", h.CreateMetric)
	repos.GET("/metrics", h.Metrics)
}
