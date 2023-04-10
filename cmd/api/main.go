package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"goulash-menu-api/internal/app"
	"goulash-menu-api/internal/config"
	"goulash-menu-api/internal/repository"
	mysql "goulash-menu-api/internal/repository/mysql"
	"goulash-menu-api/internal/service"
	http "goulash-menu-api/internal/transport/http/v1"

	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
)

func main() {
	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal(err)
	}
	cfg := config.GetConfigInstance()

	migration := flag.Bool("migration", true, "Defines the migration start option")
	flag.Parse()

	db, err := mysql.NewDB(mysql.Config{
		Username: cfg.Database.User,
		Password: cfg.Database.Password,
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		DBName:   cfg.Database.Name,
	})
	if err != nil {
		log.Fatal(err)
	}

	if *migration {
		_ = goose.SetDialect("mysql")
		if err = goose.Up(db.DB, cfg.Database.Migrations); err != nil {
			log.Fatal(err)
		}
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := http.NewHandler(services)

	srv := new(app.Server)
	go func() {
		if err = srv.Run(cfg.Http.Port, handlers.InitRoutes()); err != nil {
			log.Fatal(err)
		}
	}()
	log.Println("HTTP Сервер запущен!")

	gracefulShutdown(srv, db)
}

func gracefulShutdown(srv *app.Server, db *sqlx.DB) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("Ошибка во время остановки HTTP Сервера: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Fatalf("Ошибка во время остановки БД: %s", err.Error())
	}

	log.Println("HTTP Сервер завершил работу!")
}
