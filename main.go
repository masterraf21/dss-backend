package main

import (
	"context"
	"sync"

	"github.com/masterraf21/dss-backend/router"
)

func main() {
	ctx := context.TODO()
	handler := router.NewHandler(ctx)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		handler.HTTPStart()
	}()

	wg.Wait()
}
