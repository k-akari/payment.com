package main

import (
	"context"
	"log"
	"os"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	_, err := newConfig()
	if err != nil {
		return err
	}

	return nil
}
