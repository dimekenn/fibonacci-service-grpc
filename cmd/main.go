package main

import (
	"embed"
	"encoding/json"
	"fibonacciServiceGRPC/configs"
	"fibonacciServiceGRPC/internal/app"
	"log"
)

//go:embed configs.json
var fs embed.FS

const configName = "configs.json"

func main() {
	data, readErr := fs.ReadFile(configName)
	if readErr != nil {
		log.Fatal(readErr)
	}
	cfg := configs.NewConfig()
	if unmErr := json.Unmarshal(data, &cfg); unmErr != nil {
		log.Fatal(unmErr)
	}

	errCh := make(chan error, 1)
	go app.StartGRPCServer(errCh, cfg)

	log.Fatalf("%v", <-errCh)
}
