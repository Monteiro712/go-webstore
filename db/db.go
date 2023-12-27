package db

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)


func ConectarBancoDeDados() (*sql.DB, error) {
	// Configuração da string de conexão
	dataSourceName := "root:root@tcp(localhost:3306)/alura_loja"

	// Abre a conexão com o banco de dados
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	// Testa a conexão
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}