package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

type Produto struct {
	ID         int  
	Nome       string
	Preco      float64
	Quantidade int64
}

func main() {
	// Obtém uma conexão com o banco de dados
	db, err := ConectarBancoDeDados()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println("Conexão bem-sucedida ao banco de dados MySQL!")

	// Sempre que alguém acessar a rota "/", a função index será chamada para lidar com a requisição
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Consulta ao banco de dados para obter os produtos
		produtos := consultarProdutosDoBancoDeDados(db)

		// Executa o template e envia os produtos como dados
		temp.ExecuteTemplate(w, "Index", produtos)
	})

	// Inicia o servidor HTTP na porta 8000 e começa a lidar com as requisições
	http.ListenAndServe(":8000", nil)
}



func consultarProdutosDoBancoDeDados(db *sql.DB) []Produto {
	// Aqui você realizaria uma consulta real ao banco de dados para obter os produtos
	// Por exemplo:
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

/*package main

import (
	"html/template"
	"net/http"
)
//link with the package templates
var temp = template.Must(template.ParseGlob("templates/*.html"))

type Produto struct {
	Nome string
	Preco float64
	Quantidade int64
}

func main(){
	//sempre que alguém acessar a rota "/", a função index será chamada para lidar com a requisição
	http.HandleFunc("/", index) 
	//inicia o servidor HTTP na porta 8000 e começa a lidar com as requisições
	http.ListenAndServe(":8000", nil)
}

func index (w http.ResponseWriter, r *http.Request){
	produtos := []Produto {
		{"ps4", 2500, 50},
		{"xbox one", 2600, 62},
		{"switch", 2800, 40},   
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
*/

