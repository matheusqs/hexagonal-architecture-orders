package data

import (
	"hex-hello-world/src/domain"

	"golang.org/x/exp/slices"
)

type StaticArrayOrdersRepository struct {
	db []domain.Order
	id int
}

func (repository *StaticArrayOrdersRepository) Get(id int) domain.Order {
	index := slices.IndexFunc(repository.db, func(o domain.Order) bool { return o.Id == id })
	return repository.db[index]
}

func (repository *StaticArrayOrdersRepository) Update(id int, order domain.Order) domain.Order {
	order.Id = repository.id
	repository.db = append(repository.db, order)
	return order
}

func (repository *StaticArrayOrdersRepository) Create(order domain.Order) domain.Order {
	repository.id++
	order.Id = repository.id
	repository.db = append(repository.db, order)
	return order
}
