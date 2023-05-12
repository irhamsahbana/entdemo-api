package main

import (
	"context"
	"entdemo-api/controller"
	"entdemo-api/ent"
	"entdemo-api/repository"
	"entdemo-api/service"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", "root:@tcp(127.0.0.1:3306)/entdemo_api?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	
	ctx := context.Background()

	userRepository := repository.UserNewRepository(client, ctx)
	userService := service.UserNewService(userRepository)
	userController := controller.UserNewController(userService)

	//CreateUser(context.Background(),client)


	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/",rootController)
	r.Get("/users",userController.UserGetAllController)
	r.Post("/users",userController.UserCreateController)
	// r.Get("/cars",controller.CarController)
	// r.Get("/cars",controller.CarController)
	// r.Get("/groups",controller.GroupController)
	
    http.ListenAndServe(":3000", r)
}

func rootController(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
	return
}


