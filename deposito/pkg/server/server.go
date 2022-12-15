package server

import (
    "net/http"
    "time"

    "github.com/Racrivelari/ProjetoEquipeE/deposito/config"
    "github.com/gorilla/mux"
)

func NewServer(r *mux.Router, conf *config.Config) *http.Server {

    server := &http.Server{
        Handler:      r,
        Addr:         ": " + conf.SRV_PORT,
        WriteTimeout: 30 * time.Second,
        ReadTimeout:  30 * time.Second,
    }
    return server
}