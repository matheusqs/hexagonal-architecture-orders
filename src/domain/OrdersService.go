package domain

type Order struct {
	Id   int
	Name string
}

type OrdersService interface {
	Get(id int) Order
	Update(id int, order Order) Order
	Create(order Order) Order
}
