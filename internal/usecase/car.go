package service

import (
	"context"
	"fmt"
	"session-23/internal/data/entity"
	"session-23/internal/data/repository"
	"session-23/internal/dto"

	"time"
)

type ServiceCar struct {
	Repo repository.RepositoryCar
}

func NewServiceCar(repo *repository.RepositoryCar) *ServiceCar {
	return &ServiceCar{Repo: *repo}
}

func (s *ServiceCar) DashboardSerial(ctx context.Context, limit int) (dto.DashboardResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	cars, err := s.Repo.GetLatestCars(ctx, limit)
	if err != nil {
		return dto.DashboardResponse{}, fmt.Errorf("getLatestCars: %w", err)
	}

	total, err := s.Repo.GetTotalCars(ctx)
	if err != nil {
		return dto.DashboardResponse{}, fmt.Errorf("getTotalCars: %w", err)
	}

	stats, err := s.Repo.GetPriceStats(ctx)
	if err != nil {
		return dto.DashboardResponse{}, fmt.Errorf("getPriceStats: %w", err)
	}

	return dto.DashboardResponse{
		TotalCars: total,
		Stats:     stats,
		Cars:      cars,
	}, nil
}

func (s *ServiceCar) DashboardConcurrent(ctx context.Context, limit int) (dto.DashboardResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	carsCh := make(chan dto.ResultCars)
	totalCh := make(chan dto.ResultTotal)
	statsCh := make(chan dto.ResultStats)

	// 3 query jalan bareng
	go func() {
		cars, err := s.Repo.GetLatestCars(ctx, limit)
		carsCh <- dto.ResultCars{Data: cars, Err: err}
	}()
	go func() {
		total, err := s.Repo.GetTotalCars(ctx)
		totalCh <- dto.ResultTotal{Data: total, Err: err}
	}()
	go func() {
		stats, err := s.Repo.GetPriceStats(ctx)
		statsCh <- dto.ResultStats{Data: stats, Err: err}
	}()

	var (
		cars  []entity.Car
		total int64
		stats dto.PriceStats
	)

	for i := 0; i < 3; i++ {
		select {
		case <-ctx.Done():
			return dto.DashboardResponse{}, ctx.Err()

		case r := <-carsCh:
			if r.Err != nil {
				return dto.DashboardResponse{}, fmt.Errorf("getLatestCars: %w", r.Err)
			}
			cars = r.Data

		case r := <-totalCh:
			if r.Err != nil {
				return dto.DashboardResponse{}, fmt.Errorf("getTotalCars: %w", r.Err)
			}
			total = r.Data

		case r := <-statsCh:
			if r.Err != nil {
				return dto.DashboardResponse{}, fmt.Errorf("getPriceStats: %w", r.Err)
			}
			stats = r.Data
		}
	}

	return dto.DashboardResponse{
		TotalCars: total,
		Stats:     stats,
		Cars:      cars,
	}, nil
}
