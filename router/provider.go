package router

import "github.com/go-chi/chi"

func RegisterRouter(r *chi.Mux) {
	resgitserUserRouter(r)
}