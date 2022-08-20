package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/ats2otus/final_project/pkg/bwlist"
	"github.com/ats2otus/final_project/pkg/rate"
	"github.com/kelseyhightower/envconfig"
)

var (
	envPrefix = ""
)

func main() {
	var config Config
	if err := envconfig.Process(envPrefix, &config); err != nil {
		panic(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	service := rateService{
		blacklist: bwlist.New(),
		whitelist: bwlist.New(),

		limitByIP:     rate.NewLimiter(config.Rate.Interval, config.Rate.IP),
		limitByLogin:  rate.NewLimiter(config.Rate.Interval, config.Rate.Login),
		limitByPasswd: rate.NewLimiter(config.Rate.Interval, config.Rate.Password),
	}

	server := http.Server{
		Addr:         config.Web.Addr,
		Handler:      service.Handler(),
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
