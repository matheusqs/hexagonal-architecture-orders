package web

import (
	"encoding/json"
	"hex-hello-world/src/domain"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type OrdersController struct {
	Service domain.OrdersService
}

func (controller *OrdersController) GetOrder(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	ID, _ := strconv.Atoi(params["id"])
	order := controller.Service.Get(ID)
	json.NewEncoder(response).Encode(order)
}

func (controller *OrdersController) CreateOrder(response http.ResponseWriter, request *http.Request) {
	var order domain.Order
	_ = json.NewDecoder(request.Body).Decode(&order)
	createdOrder := controller.Service.Create(order)
	json.NewEncoder(response).Encode(createdOrder)
}

func (controller *OrdersController) UpdateOrder(response http.ResponseWriter, request *http.Request) {
}
