package models

import (
	"github.com/Monteiro712/go-webstore/db"
)

type Produto struct {
	ID         int  
	Nome       string
	Preco      float64
	Quantidade int64
}

func ConsultarProdutosDoBancoDeDados() []Produto {
	db := db.ConectarBancoDeDados()
	rows, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var produtos []Produto

	for rows.Next() {
		var produto Produto
		err := rows.Scan(&produto.ID, &produto.Nome, &produto.Preco, &produto.Quantidade)
		if err != nil {
			panic(err.Error())
		}
		produtos = append(produtos, produto)
	}
	defer db.Close()
	return produtos
}