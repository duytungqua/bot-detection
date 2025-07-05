package main

import (
	"bot-detection/config"
	"fmt"
)

func main() {
	cfg := config.Load()

	fmt.Printf("API Endpoint: %s\n", cfg.API.Endpoint)
}