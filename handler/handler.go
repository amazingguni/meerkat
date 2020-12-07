package handler

import "github.com/amazingguni/meerkat/metric"

type Handler struct {
	metricStore metric.Store
}

func NewHandler(ms metric.Store) *Handler {
	return &Handler{
		metricStore: ms,
	}
}
