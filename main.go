package main

import (
	"net/http"

	"github.com/amazingguni/meerkat/db"
	"github.com/amazingguni/meerkat/handler"
	"github.com/amazingguni/meerkat/router"
	"github.com/amazingguni/meerkat/store"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/amazingguni/meerkat/docs"
)

// @title Meerkat Swagger API
// @version 1.0
// @description Meerkat API
// @title Meerkat API

// @BasePath /api

// @schemes http https
// @produce	application/json
// @consumes application/json
func main() {
	r := router.New()
	r.GET("/swagger/*", echoSwagger.WrapHandler)
	v1 := r.Group("/api")
	d := db.New()
	db.AutoMigrate(d)
	metricStore := store.NewMetricStore(d)
	handler := handler.NewHandler(metricStore)
	handler.Register(v1)
	r.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	r.Logger.Fatal(r.Start("127.0.0.1:1323"))
}
