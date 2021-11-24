package main

import (
	"context"
	"sync"

	"github.com/masterraf21/dss-backend/router"
)

func main() {
	ctx := context.TODO()
	server := router.NewServer(ctx)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		server.Start()
	}()

	wg.Wait()
}
