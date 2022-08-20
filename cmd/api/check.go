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
func (rs *rateService) allow(w http.ResponseWriter, r *http.Request) {
	var payload CheckItem
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		rs.writeError(w, r, http.StatusBadRequest, err)
		return
	}

	ip := net.ParseIP(payload.IP)
	if rs.whitelist.Contains(ip) {
		rs.writeResult(w, r, http.StatusOK, Result{Ok: true})
		return
	}
	if rs.blacklist.Contains(ip) {
		rs.writeResult(w, r, http.StatusOK, Result{Ok: false})
		return
	}

	if ok := rs.limitByIP.Allow(payload.IP); !ok {
		rs.writeResult(w, r, http.StatusOK, Result{Ok: false})
		return
	}
	if ok := rs.limitByLogin.Allow(payload.Login); !ok {
		rs.writeResult(w, r, http.StatusOK, Result{Ok: false})
		return
	}
	if ok := rs.limitByPasswd.Allow(payload.Password); !ok {
		rs.writeResult(w, r, http.StatusOK, Result{Ok: false})
		return
	}

	rs.writeResult(w, r, http.StatusOK, Result{Ok: true})
}

// @Summary		 	Reset buckets
// @Description 	Сброс текущих ограничений по бакетам
// @Param   data 	body		ResetItem	true "данные для сброса"
// @Success	202		{object}	NoContent
// @Failure	400		{object} 	Error
// @Failure	500		{object} 	Error
// @Tags 	Common
// @Router	/reset 	[post]
func (rs *rateService) reset(w http.ResponseWriter, r *http.Request) {
	var payload ResetItem
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		rs.writeError(w, r, http.StatusBadRequest, err)
		return
	}

	rs.limitByIP.Reset(payload.IP)
	rs.limitByLogin.Reset(payload.Login)

	rs.writeResult(w, r, http.StatusAccepted, NoContent{})
}
