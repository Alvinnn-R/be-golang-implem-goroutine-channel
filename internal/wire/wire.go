package wire

import (
	"session-23/internal/adaptor"
	"session-23/internal/data/repository"
	"session-23/internal/usecase"

	"github.com/go-chi/chi/v5"
)

func Wiring(repo repository.Repository) *chi.Mux {
	router := chi.NewRouter()
	wireCar(router, repo)
	return router
}

func wireCar(router *chi.Mux, repo repository.Repository) {
	useCaseCar := usecase.NewServiceCar(&repo)
	adaptorCar := adaptor.NewAdaptorCar(useCaseCar)
	router.Get("/dashboard", adaptorCar.Dashboard)
}
