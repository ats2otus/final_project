package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/ats2otus/final_project/pkg/bwlist"
	"github.com/ats2otus/final_project/pkg/rate"
	chi "github.com/go-chi/chi/v5"
)

type rateService struct {
	blacklist bwlist.BWList
	whitelist bwlist.BWList

	limitByIP     rate.Limiter
	limitByLogin  rate.Limiter
	limitByPasswd rate.Limiter
}

// @title           Anti bruteforce
// @Description 	Сервис проверки на bruteforce
// @version         1.0.0
// @schemes			http
// @BasePath  		/v1
// @accept			json
// @produce 		json
func (rs *rateService) Handler() http.Handler {
	mux := chi.NewMux()

	mux.Route("/v1", func(r chi.Router) {

		r.Post("/allow", rs.allow)
		r.Post("/reset", rs.reset)

		r.Post("/blacklist", rs.appendBlacklist)
		r.Post("/whitelist", rs.appendWhitelist)

		r.Delete("/blacklist", rs.removeBlacklist)
		r.Delete("/whitelist", rs.removeWhitelist)
	})

	return mux
}

func (rs *rateService) writeError(w http.ResponseWriter, r *http.Request, code int, err error) {
	w.WriteHeader(code)
	if code == http.StatusNoContent || err == nil {
		return
	}
	if err := json.NewEncoder(w).Encode(Error{
		Date:  time.Now().Format(time.RFC3339),
		Error: err.Error(),
	}); err != nil {
		log.Printf("cannot write response: %v", err)
	}
}

func (rs *rateService) writeResult(w http.ResponseWriter, r *http.Request, code int, body interface{}) {
	w.WriteHeader(code)
	if code == http.StatusNoContent || body == nil {
		return
	}
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Printf("cannot write response: %v", err)
	}
}
