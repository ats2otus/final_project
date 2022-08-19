package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/kelseyhightower/envconfig"
)

var (
	envPrefix = ""
)

// @title           Anti bruteforce
// @Description 	Сервис проверки на bruteforce
// @version         1.0.0
// @schemes			http
// @BasePath  		/v1
// @accept			json
// @produce 		json

func main() {
	var config Config
	if err := envconfig.Process(envPrefix, &config); err != nil {
		panic(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	server := http.Server{
		Addr:         config.Web.Addr,
		Handler:      createHTTPHandler(),
		ReadTimeout:  config.Web.ReadTimeout,
		WriteTimeout: config.Web.WriteTimeout,
	}

	errors := make(chan error, 1)
	go func() {
		errors <- server.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			if err := server.Close(); err != nil {
				log.Fatalf("server.close: %v", err)
			}
		}
	case err := <-errors:
		log.Fatalf("server.start: %v", err)
	}
}
