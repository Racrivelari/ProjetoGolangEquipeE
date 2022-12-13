package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/Racrivelari/ProjetoEquipeE/deposito/config"
	"github.com/Racrivelari/ProjetoEquipeE/deposito/handler"
	"github.com/Racrivelari/ProjetoEquipeE/deposito/pkg/database"
	"github.com/Racrivelari/ProjetoEquipeE/deposito/pkg/service"
	"github.com/Racrivelari/ProjetoEquipeE/deposito/webui"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	default_conf := &config.Config{}

	if file_config := os.Getenv("STOQ_CONFIG"); file_config != "" {
		file, _ := os.ReadFile(file_config)
		_ = json.Unmarshal(file, &default_conf)
	}

	conf := config.NewConfig(default_conf)

	dbpool := database.NewDB(conf)
	service := service.NewProdutoService(dbpool)

	println("Driver utilizado: ", conf.DB_DRIVE)
	println(("Banco de dados: "), conf.DB_NAME)

	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	webui.RegisterUIHandlers(r, n)
	handler.RegisterAPIHandlers(r, n, service)
	http.ListenAndServe(":5000", r)                      

}


//