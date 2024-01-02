package models

import (
	"log"

	"github.com/Monteiro712/go-webstore/db"
)

type Produto struct {
	Id         int  
	Nome       string
	Preco      float64
	Quantidade int
}

func ConsultarProdutosDoBancoDeDados() []Produto {
	db := db.ConectarBancoDeDados()
	rows, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		panic(err.Error())
	}
	
	p := Produto{}
	produtos := []Produto{}

	for rows.Next() {
		var Id int
		var Nome string
		var Preco float64
		var Quantidade int

		err := rows.Scan(&Id, &Nome, &Preco, &Quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = Id
		p.Nome = Nome
		p.Preco = Preco
		p.Quantidade = Quantidade
		
		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func CriarNovoProduto(nome string, preco float64, quantidade int){
	db.ConectarBancoDeDados()

	insereDadosNoBanco, err := db.ConectarBancoDeDados().Prepare("INSERT INTO produtos(nome, preco, quantidade) VALUES(?, ?, ?)")


	if err != nil{
		panic(err.Error())
	}
	insereDadosNoBanco.Exec(nome, preco, quantidade)
	defer db.ConectarBancoDeDados().Close()
}

func DeletaProduto(id string){
	db := db.ConectarBancoDeDados()

	deletarOProduto, err := db.Prepare("delete from produtos where id=?")
	if err != nil{
		panic(err.Error())
	}
	deletarOProduto.Exec(id)
	defer db.Close() 
}

func EditaProduto(id string) Produto {
	db := db.ConectarBancoDeDados()

	produtoDoBanco, err := db.Query("select * from produtos where id=?", id)
	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for produtoDoBanco.Next() {
		var id int
		var quantidade int
		var preco float64
		var nome string
		err = produtoDoBanco.Scan(&id, &nome, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}
	defer db.Close()
	return produtoParaAtualizar
}

func AtualizaProduto(id int, nome string, preco float64, quantidade int) {
    db := db.ConectarBancoDeDados()

    _, err := db.Exec("UPDATE produtos SET nome=?, preco=?, quantidade=? WHERE id=?", nome, preco, quantidade, id)
    if err != nil {
        log.Println("Erro na atualização do produto:", err)
    }

    defer db.Close()
}
