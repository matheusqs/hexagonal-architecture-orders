package domain

type OrdersServiceImpl struct {
	Orders Orders
}

func (o *OrdersServiceImpl) Get(id int) Order {
	return o.Orders.Get(id)
}

func (o *OrdersServiceImpl) Update(id int, order Order) Order {
	return o.Orders.Update(id, order)
}

func (o *OrdersServiceImpl) Create(order Order) Order {
	return o.Orders.Create(order)
}
