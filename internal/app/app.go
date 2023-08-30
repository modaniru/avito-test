package app

import (
	"context"
	"database/sql"
	"fmt"
	log "log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
	"github.com/modaniru/avito/internal/controller"
	"github.com/modaniru/avito/internal/service"
	"github.com/modaniru/avito/internal/storage"
	yandexdrive "github.com/modaniru/avito/internal/yandex_drive"
)

func App() {
	config := configure()
	configureLogger(config.LogLevel)
	log.Debug("logger was successfully configured")

	log.Debug("open database connection...")
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", config.DatabaseUser, config.DatabasePassword, config.DatabaseHost, config.DatabaseName))
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
	yandex, err := yandexdrive.NewYandexDisk(config.YandexToken)
	if err != nil {
		log.Error("create yandex disk error", log.String("error", err.Error()))
		return
	}
	if yandex.IsAvailible() {
		log.Debug("history service enable")
	} else {
		log.Debug("history service disable")
	}
	scheduler := service.NewScheduler(storage.Follow)
	service := service.NewService(storage, yandex)
	r := chi.NewRouter()
	controller.NewRouter(r, service)
	srv := http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%s", config.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Debug("finish DI")
	log.Info("run scheduler...")
	channel := scheduler.RunScheduler()
	go func() {
		log.Info("run server in " + config.Port + " port...")
		err := srv.ListenAndServe()
		if err != nil {
			log.Error("error", log.String("error", err.Error()))
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-quit:
		log.Info("shutdown server")
		srv.Shutdown(context.Background())
	case <-channel:
		srv.Shutdown(context.Background())
		log.Error("scheduler error")
	}
}
