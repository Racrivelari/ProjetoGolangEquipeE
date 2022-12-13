package main

import (
	"encoding/json"
	"fmt"

	// "log"
	"net/http"
	"os"

	"github.com/Racrivelari/ProjetoGolangOFICIAL/deposito/config"
	"github.com/Racrivelari/ProjetoGolangOFICIAL/deposito/handler"
	"github.com/Racrivelari/ProjetoGolangOFICIAL/deposito/pkg/database"
	"github.com/Racrivelari/ProjetoGolangOFICIAL/deposito/pkg/service"
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

	r.HandleFunc("/", redirect)
	handler.RegisterAPIHandlers(r, n, service)
	fmt.Println("Escutando na porta 5000")
	//log.Fatal(http.ListenAndServe(":5000", r))

	fs := http.FileServer(http.Dir("webui/dist/spa"))
	r.Handle("/webui/", http.StripPrefix("/webui/", fs)) //aq vc passa o r inves do http.handle, pq vc criou suas proprias rotas no router, ou seja vc n ta usando padrao go e sim um proprio
	http.ListenAndServe(":5000", r)                      //tem q ser r, eu só passaria nil aq, se eu n tivesse feito minhas rotas proprias

	//DEU CERTO COM ESSE JEITO DO ISAEL AQ, funcionou o frontend, e tipo, eu tento logar ele da uma msg verde de "bem vindo admin" porem n redireciona pra home n sei pq
	//outro ponto, eu tirei a main da pasta raiz e joguei pra cmd dnv, ele voltou a funcionar nao faco ideia pq KKKKKKKKKKKKKKK

	//VC PODE PULAR A TELA DE LOGIN, VC COLOCA O ENDERECO/PRODUCTS, dai pula, mas n funciona o front, porem ele puxa a lista de produtos

}

func redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/webui/", http.StatusMovedPermanently)
}
