package handler

import (
	"github.com/amazingguni/meerkat/model"
	"github.com/labstack/echo/v4"
)

type metricCreateRequest struct {
	Metric struct {
		Value *float32 `json:"value" validate:"required"`
	}
}

func (r *metricCreateRequest) bind(c echo.Context, m *model.Metric) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	m.Value = *r.Metric.Value
	return nil
}
