package config

import (
	"database/sql"
	"net/http"
	"tidy/pkg/handler"
	"tidy/pkg/repository"
	"tidy/pkg/service"
)

func Config(db *sql.DB) *http.ServeMux {
	router := http.NewServeMux()
	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	handler.Register(router)

	return router
}
