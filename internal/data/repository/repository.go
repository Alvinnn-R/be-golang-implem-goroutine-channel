package repository

import "session-23/internal/usecase"


type Repository struct {
	RepositoryCar *RepositoryCar
}

func NewRepository(usecase *usecase.ServiceCar) *Repository {
	return &Repository{
		RepositoryCar: NewRepositoryCar(usecase),
	}
}