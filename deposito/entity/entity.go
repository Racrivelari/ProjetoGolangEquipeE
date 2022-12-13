package entity

import (
	"encoding/json"
	"log"
)

type ProdutoInterface interface {
	String() string
}

type Product struct { //NOME DE VARIAVEL TEM Q SER MAISUCULA POR CAUSA DO JSON
	ID        int `json:"id"`
	Name      string `json:"name"`
	Price     float64 `json:"price"`
	Code      string  `json:"code"`
	CreatedAt string `json:"created_at,omitempty"`
}

func (p Product) String() string {
	data, err := json.Marshal(p)

	if err != nil {
		log.Println("erro ao converter produto p json")
		return ""
	}

	return string(data)
}

type ProdutoList struct {
	List []*Product `json:"list"`
}

func (pl ProdutoList) String() string {
	data, err := json.Marshal(pl)

	if err != nil {
		log.Println("erro ao converter lista de produto p json")
		log.Println(err.Error())
		return ""
	}

	return string(data)
}

func NewProduto(nome string, price float64, code string,) *Product {
	return &Product{
		Name:  nome,
		Price: price,
		Code:  code,
	}
}




