package app

import (
	"database/sql"
	"fmt"
	log "log/slog"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
	"github.com/modaniru/avito/internal/controller"
	"github.com/modaniru/avito/internal/service"
	"github.com/modaniru/avito/internal/storage"
)

func App() {
	config := configure()
	configureLogger(config.LogLevel)
	log.Debug("logger was successfully configured")

	log.Debug("open database connection...")
	db, err := sql.Open("postgres", config.DatabaseSource)
	if err != nil {
		log.Error("open database error", log.String("error", err.Error()))
		return
	}
	err = db.Ping()
	if err != nil {
		log.Error("ping database error", log.String("error", err.Error()))
		return
	}
	log.Debug("database successfully load")

	log.Debug("start DI...")
	storage := storage.NewStorage(db)
	service := service.NewService(storage)
	r := chi.NewRouter()
	controller.NewRouter(r, service)
	log.Debug("finish DI")

	http.ListenAndServe(fmt.Sprintf(":%s", config.Port), r)

}
