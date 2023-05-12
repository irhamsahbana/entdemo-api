package router

import (
	bootstrap "entdemo-api/bootsrap"
	"entdemo-api/controller"
	"entdemo-api/repository"
	"entdemo-api/service"

	"github.com/go-chi/chi"
)

func resgitserUserRouter(r *chi.Mux) {
	client := bootstrap.App.Ent
	userRepository := repository.UserNewRepository(client)
	userService := service.UserNewService(userRepository)
	userController := controller.UserNewController(userService)

	r.Route("/api/v1",func(r chi.Router) {
		r.Get("/users",userController.UserGetAllController)
		r.Get("/users/{id}",userController.UserGetByIDController)
		r.Post("/users",userController.UserCreateController)
		
	})


}