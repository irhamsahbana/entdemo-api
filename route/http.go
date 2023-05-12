package route

import (
	"entdemo-api/bootstrap"
	"entdemo-api/controller"
	"entdemo-api/repository"
	"entdemo-api/service"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func NewHttpRoute(r *chi.Mux) {
	client := bootstrap.App.Ent

	userRepository := repository.UserNewRepository(client)
	userService := service.UserNewService(userRepository)
	userController := controller.UserNewController(userService)

	r.Use(middleware.Logger)
	// use prefix
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/", rootController)
		r.Get("/users", userController.UserGetAllController)
		r.Post("/users", userController.UserCreateController)
	})
	// r.Get("/cars",controller.CarController)
	// r.Get("/cars",controller.CarController)
	// r.Get("/groups",controller.GroupController)

	http.ListenAndServe(":3000", r)
}

func rootController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
	return
}
