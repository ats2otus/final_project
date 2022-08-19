package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func createHTTPHandler() http.Handler {
	mux := chi.NewMux()

	mux.Route("/v1", func(r chi.Router) {
		rg := rateGroup{}

		r.Post("/allow", rg.allow)
		r.Post("/reset", rg.reset)

		r.Post("/blacklist", rg.appendBlacklist)
		r.Post("/whitelist", rg.appendWhitelist)

		r.Delete("/blacklist", rg.removeBlacklist)
		r.Delete("/whitelist", rg.removeWhitelist)
	})

	return mux
}

type rateGroup struct {
}

// @Summary		 	Bruteforce detection
// @Description 	Проверка на bruteforce
// @Param   data 	body		CheckItem	true "данные для проверки"
// @Success	200		{object}	Result
// @Failure	400		{object} 	Error
// @Failure	500		{object} 	Error
// @Tags 	Common
// @Router	/allow 	[post]
func (rg *rateGroup) allow(w http.ResponseWriter, r *http.Request) {
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
	w.WriteHeader(http.StatusNotImplemented)
}

// @Summary		 	Blacklist append
// @Description 	Добавление IP в blacklist
// @Param   data 	body		ListItem	true "подсеть (IP + маска)"
// @Success	202		{object}	NoContent
// @Failure	400		{object} 	Error
// @Failure	500		{object} 	Error
// @Tags 	BlackList
// @Router	/blacklist 	[post]
func (rg *rateGroup) appendBlacklist(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
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
	w.WriteHeader(http.StatusNotImplemented)
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
	w.WriteHeader(http.StatusNotImplemented)
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
	w.WriteHeader(http.StatusNotImplemented)
}
