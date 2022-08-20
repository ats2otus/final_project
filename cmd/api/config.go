package main

import "time"

type Config struct {
	Web struct {
		Addr         string        `default:":8080"`
		ReadTimeout  time.Duration `default:"5s"`
		WriteTimeout time.Duration `default:"5s"`
	}
	Rate struct {
		IP       int           `default:"1000"`
		Login    int           `default:"10"`
		Password int           `default:"100"`
		Interval time.Duration `default:"1m"`
	}
}
