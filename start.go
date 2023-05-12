package main

import (
	"entdemo-api/router"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	router.RegisterRouter(r)
	
    http.ListenAndServe(":3000", r)
}

func rootController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
	return
}


