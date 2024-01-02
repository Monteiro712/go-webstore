package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
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
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Delete(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        idDoProduto := r.FormValue("id")
        models.DeletaProduto(idDoProduto)
    }
    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Edit(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    idDoProduto := vars["id"]
    produto := models.EditaProduto(idDoProduto)
    temp.ExecuteTemplate(w, "Edit", produto)
}


func Update(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    nome := r.FormValue("nome")
    preco := r.FormValue("preco")
    quantidade := r.FormValue("quantidade")

    // Converte os valores conforme necessário (ex: preço para float, quantidade para int)

    // Validação dos campos e manipulação de erros

    // Atualiza o produto no banco de dados usando os valores recebidos
    idConvertidaParaInt, err := strconv.Atoi(id)
    if err != nil {
        log.Println("Erro na conversão do ID para int:", err)
        // Trate o erro de forma apropriada, como retornar um código de status HTTP 400 (Bad Request)
        http.Error(w, "Erro na conversão do ID", http.StatusBadRequest)
        return
    }

    precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
    if err != nil {
        log.Println("Erro na conversão do preço para float64:", err)
        http.Error(w, "Erro na conversão do preço", http.StatusBadRequest)
        return
    }

    quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
    if err != nil {
        log.Println("Erro na conversão da quantidade para int:", err)
        http.Error(w, "Erro na conversão da quantidade", http.StatusBadRequest)
        return
    }

    // Atualiza o produto no banco de dados
    models.AtualizaProduto(idConvertidaParaInt, nome, precoConvertidoParaFloat, quantidadeConvertidaParaInt)

    // Redireciona ou responde de acordo com a lógica do seu aplicativo
    http.Redirect(w, r, "/", http.StatusSeeOther)
}


