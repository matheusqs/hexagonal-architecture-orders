package domain

type Orders interface {
	Get(id int) Order
	Update(id int, order Order) Order
	Create(order Order) Order
}
