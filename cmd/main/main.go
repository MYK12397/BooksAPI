package main

import (
	"log"
	"net/http"
	"os"

	"github.com/MYK12397/BooksAPI/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {

		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("HOST")
	r := mux.NewRouter()
	routes.BookRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(host, r))

}
