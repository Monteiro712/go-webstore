package models

import "github.com/Monteiro712/db"

type Produto struct {
	ID         int  
	Nome       string
	Preco      float64
	Quantidade int64
}

func consultarProdutosDoBancoDeDados(db *sql.DB) []Produto {
	db := db.
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

	return produtos
}