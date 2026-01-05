package main

import (
	"session-23/cmd"
	"session-23/internal/data/repository"
	"session-23/internal/wire"
	"session-23/pkg/database"
	"session-23/pkg/utils"
	"time"
)

type Item struct {
	Name  string
	Price int
	Qty   int
}

type Result struct {
	Subtotal int
}

var cartData = []Item{
	{Name: "Beras 5kg", Price: 75000, Qty: 1},
	{Name: "Gula 1kg", Price: 18000, Qty: 2},
	{Name: "Minyak 2L", Price: 32000, Qty: 1},
}

func main() {
	config, err := utils.ReadConfiguration()
	if err != nil {
	}
	db, err := database.InitDB(config.DB)

	if err != nil {
	}
	repo := repository.NewRepository(db)
	router := wire.Wiring(repo)
	cmd.APiserver(router)

}

func CalculateTotalItem(cart []Item) int {
	total := 0
	for _, it := range cart {
		time.Sleep(200 * time.Microsecond)
		subtotal := it.Price * it.Qty
		total += subtotal
	}
	return total
}

func CalculateTotalItemConcurrent(cart []Item) int {
	resultCh := make(chan Result)

	for _, it := range cart {
		go func(item Item) {
			time.Sleep(200 * time.Microsecond)
			subtotal := item.Price * item.Qty
			resultCh <- Result{Subtotal: subtotal}
		}(it)
	}

	total := 0
	for i := 0; i < len(cart); i++ {
		res := <-resultCh
		total += res.Subtotal
	}

	return total
}
