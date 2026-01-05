package adaptor

import (
	"net/http"
	"session-23/internal/usecase"
)

type AdaptorCar struct {
	Usecase usecase.ServiceCar
}

func NewAdaptorCar(usecase *usecase.ServiceCar) *AdaptorCar {
	return &AdaptorCar{Usecase: *usecase}
}

func (usecaseAdaptor *AdaptorCar) Dashboard(w http.ResponseWriter, r *http.Request) {

}
