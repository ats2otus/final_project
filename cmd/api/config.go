package main

import "time"

type Config struct {
	Web struct {
		Addr         string        `default:":8080"`
		ReadTimeout  time.Duration `default:"5s"`
		WriteTimeout time.Duration `default:"5s"`
	}
	Rate struct {
		IP       string `default:"1000"`
		Login    string `default:"10"`
		Password string `default:"100"`
	}
}
