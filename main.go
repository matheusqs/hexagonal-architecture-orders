package main

import (
	"fmt"
	"hex-hello-world/src/data"
	"hex-hello-world/src/domain"
	"hex-hello-world/src/web"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// função principal
func main() {
	ordersRepository := &data.StaticArrayOrdersRepository{}

	ordersService := &domain.OrdersServiceImpl{
		Orders: ordersRepository,
	}

	ordersController := web.OrdersController{
		Service: ordersService,
	}

	router := mux.NewRouter()
	router.HandleFunc("/order/{id}", ordersController.GetOrder).Methods("GET")
	router.HandleFunc("/order", ordersController.CreateOrder).Methods("POST")
	router.HandleFunc("/order/{id}", ordersController.UpdateOrder).Methods("PUT")

	fmt.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":8000", router))
}
