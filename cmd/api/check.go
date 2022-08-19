package main

import (
	"encoding/json"
	"net"
	"net/http"
)

// @Summary		 	Bruteforce detection
// @Description 	Проверка на bruteforce
// @Param   data 	body		CheckItem	true "данные для проверки"
// @Success	200		{object}	Result
// @Failure	400		{object} 	Error
// @Failure	500		{object} 	Error
// @Tags 	Common
// @Router	/allow 	[post]
func (rg *rateGroup) allow(w http.ResponseWriter, r *http.Request) {
	var payload CheckItem
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		rg.writeError(w, r, http.StatusBadRequest, err)
		return
	}

	if payload.IP != "" {
		ip := net.ParseIP(payload.IP)
		if rg.whitelist.Contains(ip) {
			rg.writeResult(w, r, http.StatusOK, Result{Ok: true})
			return
		}
		if rg.blacklist.Contains(ip) {
			rg.writeResult(w, r, http.StatusOK, Result{Ok: false})
			return
		}
	}

	w.WriteHeader(http.StatusNotImplemented)
}

// @Summary		 	Reset buckets
// @Description 	Сброс текущих ограничений по бакетам
// @Param   data 	body		ResetItem	true "данные для сброса"
// @Success	202		{object}	NoContent
// @Failure	400		{object} 	Error
// @Failure	500		{object} 	Error
// @Tags 	Common
// @Router	/reset 	[post]
func (rg *rateGroup) reset(w http.ResponseWriter, r *http.Request) {
	var payload ResetItem
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		rg.writeError(w, r, http.StatusBadRequest, err)
		return
	}
	w.WriteHeader(http.StatusNotImplemented)
}
