package main

import (
	"net/http"
	"github.com/Monteiro712/go-webstore/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}



