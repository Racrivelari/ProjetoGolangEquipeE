package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/Racrivelari/ProjetoEquipeE/deposito/entity"
	"github.com/Racrivelari/ProjetoEquipeE/deposito/pkg/service"
	"github.com/gorilla/mux"
)

func getAllProduct(service service.ProdutoServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		all := service.GetAll()
		err := json.NewEncoder(w).Encode(all)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(all.String()))
			return
		}
	})
}

func getProduct(service service.ProdutoServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)

		ID, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"ERRO, id necessário"}`))
			return
		}

		produto := service.GetByID(&ID)
		if produto.ID == 0 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"ERRO, produto não encontrado"}`))
			return
		}

		err = json.NewEncoder(w).Encode(produto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"ERRO, conversão de produto em json falhou"}`))
			return
		}
	})
}

func createProduct(service service.ProdutoServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		produto := entity.Product{}

		err := json.NewDecoder(r.Body).Decode(&produto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"ERRO, conversão de produto em json falhou"}`))
			return
		}

		last_id := service.Create(&produto)
		if last_id == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"ERRO, inserção de produto falhou"}`))
			return
		}

		produto = *service.GetByID(&last_id)

		err = json.NewEncoder(w).Encode(produto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"ERRO, conversão de produto em json falhou"}`))
			return
		}
	})
}

func updateProduct(service service.ProdutoServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)

		ID, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"ERRO, id necessário"}`))
			return
		}

		produto := entity.Product{}

		err = json.NewDecoder(r.Body).Decode(&produto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"ERRO, conversão de produto em json falhou"}`))
			return
		}

		rows_affected := service.Update(&ID, &produto)
		if rows_affected == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"ERRO, atualização do produto falhou"}`))
			return
		}

		produto = *service.GetByID(&ID)

		err = json.NewEncoder(w).Encode(produto)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"ERRO, conversão de produto em json falhou"}`))
			return
		}
	})
}

func deleteProduct(service service.ProdutoServiceInterface) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)

		ID, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"ERRO, é necessário um ID"}`))
			return
		}

		rows_affected := service.Delete(&ID)
		if rows_affected == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"ERRO, deleção de produto falhou"}`))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Produto deletado com sucesso"}`))
	})
}
