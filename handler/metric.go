package handler

import (
	"net/http"
	"strconv"

	"github.com/amazingguni/meerkat/model"
	"github.com/amazingguni/meerkat/utils"
	"github.com/labstack/echo/v4"
)

// CreateMetric godoc
// @Summary Create a metric
// @Description Create a metric. Auth is required
// @ID create-metric
// @Tags metric
// @Accept  json
// @Produce  json
// @Param metric body metricCreateRequest true "Metric to create"
// @Success 201 {object} metricResponse
// @Failure 401 {object} utils.Error
// @Failure 422 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Security ApiKeyAuth
// @Router /repos/metrics [post]
func (h *Handler) CreateMetric(c echo.Context) error {
	var m model.Metric
	req := &metricCreateRequest{}
	if err := req.bind(c, &m); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	err := h.metricStore.CreateMetric(&m)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusCreated, newMetricResponse(c, &m))
}

// Metrics godoc
// @Summary Get recent metrics globally
// @Description Get most recent metrics globally. Use query parameters to filter results. Auth is optional
// @ID get-metrics
// @Tags metric
// @Accept  json
// @Produce  json
// @Param limit query integer false "Limit number of articles returned (default is 20)"
// @Param offset query integer false "Offset/skip number of articles (default is 0)"
// @Success 200 {object} metricListResponse
// @Failure 500 {object} utils.Error
// @Router /repos/metrics [get]
func (h *Handler) Metrics(c echo.Context) error {
	var (
		metrics []model.Metric
		count   int
	)
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 20
	}
	metrics, count, err = h.metricStore.List(offset, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}
	return c.JSON(http.StatusOK, newMetricListResponse(c, metrics, count))
}
