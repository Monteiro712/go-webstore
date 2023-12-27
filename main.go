package main

import (
	"html/template"
	"net/http"
	"github.com/Monteiro712/go-webstore/models"
)
	

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	http.HandleFunc("/", index)

	http.ListenAndServe(":8000", nil)
}

func index (w http.ResponseWriter, r *http.Request){
	todosOsProdutos := models.ConsultarProdutosDoBancoDeDados()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}


