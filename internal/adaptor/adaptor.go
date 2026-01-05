package adaptor

import "session-23/internal/usecase"

type AdaptorCar struct {
	usecase.ServiceCar
}

func NewAdaptorCar(usecase *usecase.ServiceCar) *AdaptorCar {
	return &AdaptorCar{ServiceCar: *usecase}
}

func (a *AdaptorCar) Service() *usecase.ServiceCar {
	return &a.ServiceCar
}