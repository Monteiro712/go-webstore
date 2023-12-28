package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/Monteiro712/go-webstore/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index (w http.ResponseWriter, r *http.Request){
	todosOsProdutos := models.ConsultarProdutosDoBancoDeDados()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New (w http.ResponseWriter, r *http.Request){
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert (w http.ResponseWriter, r *http.Request){

	if r.Method == "POST"{
		nome := r.FormValue("nome")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)

		if err != nil{
			log.Println("erro ao converter preço:", err)
		}

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		
		if err != nil{
			log.Println("erro ao converter qtde:", err)
		}
		models.CriarNovoProduto(nome, precoConvertidoParaFloat, quantidadeConvertidaParaInt)
	}
	http.Redirect(w, r, "/", 301)
}