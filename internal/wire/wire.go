package wire

import (
	"session-23/internal/adaptor"
	"session-23/internal/usecase"

	"github.com/go-chi/chi/v5"
)

func Wiring() *chi.Mux {
	r := chi.NewRouter()

	return r
}

func wireCar(r *chi.Mux) {
	useCaseCar := usecase.NewServiceCar(&repo)
	adaptorCar := adaptor.NewAdaptorCar(useCaseCar)
	r.Get("/dashboard", adaptorCar.Dashboard)
}
