package handler

import (
	"net/http"
	"github.com/urfave/negroni"
)

func applicationJSON() negroni.Handler {	//converte p json
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next(w, r)
	})
}


