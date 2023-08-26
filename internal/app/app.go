package app

import (
	"database/sql"
	log "log/slog"

	_ "github.com/lib/pq"
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
	_ = storage.NewStorage(db)
	log.Debug("finish DI")
}
