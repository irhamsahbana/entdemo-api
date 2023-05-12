package main

import (
	"entdemo-api/route"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := chi.NewRouter()
	route.NewHttpRoute(r)
}
