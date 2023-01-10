package main

import (
	"log"
	"net/http"

	"github.com/akash/bookApi/pkg/routes"
	"github.com/gorilla/mux"
	_ "gorm.io/driver/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.BookRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":9010", r))
}
