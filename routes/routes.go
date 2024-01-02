package routes

import (
    "net/http"
    "github.com/gorilla/mux"
    "github.com/Monteiro712/go-webstore/controllers"
)

func CarregaRotas() {
    r := mux.NewRouter()

    r.HandleFunc("/", controllers.Index)
    r.HandleFunc("/new", controllers.New)
    r.HandleFunc("/insert", controllers.Insert)
    r.HandleFunc("/delete", controllers.Delete)
    r.HandleFunc("/edit", controllers.Edit)
    r.HandleFunc("/update", controllers.Update).Methods("PUT")

    http.Handle("/", r)
}
