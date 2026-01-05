package dto

import "session-23/internal/data/entitiy"

type PriceStats struct {
	Min float64
	Max float64
	Avg float64
}

type DashboardResponse struct {
	TotalCars int64
	Stats     PriceStats
	Cars      []entitiy.Car
}

type ResultCars struct {
	Data []entitiy.Car
	Err  error
}
type ResultTotal struct {
	Data int64
	Err  error
}
type ResultStats struct {
	Data PriceStats
	Err  error
}
