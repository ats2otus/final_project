package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/crayoned/anti_bruteforce/pkg/bwlist"
)

type rateGroup struct {
	blacklist bwlist.BWList
	whitelist bwlist.BWList
}

func (rg *rateGroup) writeError(w http.ResponseWriter, r *http.Request, code int, err error) {
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

func (rg *rateGroup) writeResult(w http.ResponseWriter, r *http.Request, code int, body interface{}) {
	w.WriteHeader(code)
	if code == http.StatusNoContent || body == nil {
		return
	}
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Printf("cannot write response: %v", err)
	}
}
