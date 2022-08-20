package main

import (
	"net/http"

	"github.com/ats2otus/final_project/pkg/bwlist"
	"github.com/go-chi/chi/v5"
)

// @title           Anti bruteforce
// @Description 	Сервис проверки на bruteforce
// @version         1.0.0
// @schemes			http
// @BasePath  		/v1
// @accept			json
// @produce 		json
func createHTTPHandler() http.Handler {
	mux := chi.NewMux()

	mux.Route("/v1", func(r chi.Router) {
		rg := rateGroup{
			blacklist: bwlist.New(),
			whitelist: bwlist.New(),
		}

		r.Post("/allow", rg.allow)
		r.Post("/reset", rg.reset)

		r.Post("/blacklist", rg.appendBlacklist)
		r.Post("/whitelist", rg.appendWhitelist)

		r.Delete("/blacklist", rg.removeBlacklist)
		r.Delete("/whitelist", rg.removeWhitelist)
	})

	return mux
}
