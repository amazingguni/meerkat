package handler

import (
	"time"

	"github.com/amazingguni/meerkat/model"
	"github.com/labstack/echo/v4"
)

type metricResponse struct {
	Value     float32   `json:"value"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type metricListResponse struct {
	Metrics []*metricResponse `json:"metrics"`
	Count   int               `json:"count"`
}

func newMetricResponse(c echo.Context, m *model.Metric) *metricResponse {
	mr := new(metricResponse)
	mr.Value = m.Value
	mr.CreatedAt = m.CreatedAt
	mr.UpdatedAt = m.UpdatedAt
	return mr
}

func newMetricListResponse(c echo.Context, metrics []model.Metric, count int) *metricListResponse {
	r := new(metricListResponse)
	r.Metrics = make([]*metricResponse, 0)
	for _, m := range metrics {
		mr := new(metricResponse)
		mr.Value = m.Value
		mr.CreatedAt = m.CreatedAt
		mr.UpdatedAt = m.UpdatedAt
		r.Metrics = append(r.Metrics, mr)
	}
	return r
}
