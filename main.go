package main

import (
	"fmt"
	"log"
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
	// Read configuration from .env
	config, err := utils.ReadConfiguration()
	if err != nil {
		log.Fatal("Error reading configuration: ", err)
	}
	fmt.Printf("Configuration loaded: %s\n", config.AppName)

	// Initialize logger
	logger, err := utils.InitLogger(config.PathLogging, config.Debug)
	if err != nil {
		log.Fatal("Error initializing logger: ", err)
	}
	defer logger.Sync()
	fmt.Println("Logger initialized")

	// Initialize database connection
	db, err := database.InitDB(config.DB)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	fmt.Println("Database connected successfully")

	// Initialize repository
	repo := repository.NewRepository(db)
	router := wire.Wiring(repo, config, logger)
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
