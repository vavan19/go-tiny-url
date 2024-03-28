package main

import (
	// "fmt"
	// "log"
	// "net/http"
	"fmt"
	"net/http"
	"tiny_url/internal/config"
	"tiny_url/internal/counter"
	"tiny_url/internal/handler"
	"tiny_url/internal/storage"
)

func main() {
	cfg := config.Load()
	if err := cfg.Validate(); err != nil {
		panic(err)
	}

	storage := storage.NewMongoURLStore(cfg.MongoDBURL, "tinyurl", "urls")

	counter := counter.New(cfg.StartRange)

	urlHandler := handler.NewURLHandler(storage, counter)


	http.HandleFunc("/shorten", urlHandler.ShortenURL)
	http.HandleFunc("/", urlHandler.Redirect)

	fmt.Printf("Server listening on port 8080 with range %d to %d...\n", cfg.StartRange, cfg.EndRange)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
