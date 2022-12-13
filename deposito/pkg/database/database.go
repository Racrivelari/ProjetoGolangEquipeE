package database

import (
	"database/sql"
	"github.com/Racrivelari/ProjetoGolangOFICIAL/deposito/config"
)

type DatabaseInterface interface {
	GetDB() (DB *sql.DB)
	Close() error
}

type dabase_pool struct {
	DB *sql.DB
}

var dbpool = &dabase_pool{}

func NewDB(conf *config.Config) *dabase_pool {
	dbpool = MySQLConn(conf)
	return dbpool
}

func (d *dabase_pool) Close() error {
	err := d.DB.Close()
	if err != nil {
		return err
	}
	dbpool = &dabase_pool{}
	return err
}

func (d *dabase_pool) GetDB() (DB *sql.DB) {
	return d.DB
}
