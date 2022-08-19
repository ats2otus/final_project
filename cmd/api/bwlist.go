package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

// @Summary		 	Blacklist append
// @Description 	Добавление IP в blacklist
// @Param   data 	body		ListItem	true "подсеть (IP + маска)"
// @Success	202		{object}	NoContent
// @Failure	400		{object} 	Error
// @Failure	500		{object} 	Error
// @Tags 	BlackList
// @Router	/blacklist 	[post]
func (rg *rateGroup) appendBlacklist(w http.ResponseWriter, r *http.Request) {
	var payload ListItem
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		rg.writeError(w, r, http.StatusBadRequest, err)
		return
	}
	_, subnet, err := net.ParseCIDR(payload.Subnet)
	if err != nil {
		rg.writeError(w, r, http.StatusBadRequest, err)
		return
	}

	rg.blacklist.Append(subnet)
	rg.writeResult(w, r, http.StatusAccepted, NoContent{})
}

// @Summary		 	Blacklist remove
// @Description 	Удаление IP из blacklist
// @Param   subnet	query		string		true	"подсеть (IP + маска)"
// @Success	202		{object}	NoContent
// @Failure	400		{object} 	Error
// @Failure	500		{object} 	Error
// @Tags 	BlackList
// @Router	/blacklist 	[delete]
func (rg *rateGroup) removeBlacklist(w http.ResponseWriter, r *http.Request) {
	payload := r.URL.Query().Get("subnet")
	if payload == "" {
		rg.writeError(w, r, http.StatusBadRequest, fmt.Errorf("missing subnet"))
		return
	}
	_, subnet, err := net.ParseCIDR(payload)
	if err != nil {
		rg.writeError(w, r, http.StatusBadRequest, err)
		return
	}

	rg.blacklist.Remove(subnet)
	rg.writeResult(w, r, http.StatusAccepted, NoContent{})
}

// @Summary		 	Whitelist append
// @Description 	Добавление IP в whitelist
// @Param   data 	body		ListItem	true "подсеть (IP + маска)"
// @Success	202		{object}	NoContent
// @Failure	400		{object} 	Error
// @Failure	500		{object} 	Error
// @Tags 	WhiteList
// @Router	/whitelist 	[post]
func (rg *rateGroup) appendWhitelist(w http.ResponseWriter, r *http.Request) {
	var payload ListItem
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		rg.writeError(w, r, http.StatusBadRequest, err)
		return
	}
	_, subnet, err := net.ParseCIDR(payload.Subnet)
	if err != nil {
		rg.writeError(w, r, http.StatusBadRequest, err)
		return
	}

	rg.whitelist.Append(subnet)
	rg.writeResult(w, r, http.StatusAccepted, NoContent{})
}

// @Summary		 	Whitelist remove
// @Description 	Удаление IP из whitelist
// @Param   subnet	query		string		true	"подсеть (IP + маска)"
// @Success	202		{object}	NoContent
// @Failure	400		{object} 	Error
// @Failure	500		{object} 	Error
// @Tags 	WhiteList
// @Router	/whitelist 	[delete]
func (rg *rateGroup) removeWhitelist(w http.ResponseWriter, r *http.Request) {
	payload := r.URL.Query().Get("subnet")
	if payload == "" {
		rg.writeError(w, r, http.StatusBadRequest, fmt.Errorf("missing subnet"))
		return
	}
	_, subnet, err := net.ParseCIDR(payload)
	if err != nil {
		rg.writeError(w, r, http.StatusBadRequest, err)
		return
	}

	rg.whitelist.Remove(subnet)
	rg.writeResult(w, r, http.StatusAccepted, NoContent{})
}
