package handler

import (
	"github.com/Racrivelari/ProjetoEquipeE/deposito/pkg/service"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func RegisterAPIHandlers(r *mux.Router, n *negroni.Negroni, service service.ProdutoServiceInterface) {

	s := r.PathPrefix("/api/v1").Subrouter()
	n.Use(applicationJSON())

	s.Handle("/user/login", n.With(
	)).Methods("POST", "OPTIONS")

	s.Handle("/products", n.With(
		negroni.Wrap(getAllProduct(service)),
	)).Methods("GET", "OPTIONS")

	s.Handle("/product/{id}", n.With(
		negroni.Wrap(getProduct(service)),
	)).Methods("GET", "OPTIONS")

	s.Handle("/product", n.With(
		negroni.Wrap(createProduct(service)),
	)).Methods("POST", "OPTIONS")

	s.Handle("/product/{id}", n.With(
		negroni.Wrap(updateProduct(service)),
	)).Methods("PUT", "OPTIONS")

	s.Handle("/product/{id}", n.With(
		negroni.Wrap(deleteProduct(service)),
	)).Methods("DELETE", "OPTIONS")

}
