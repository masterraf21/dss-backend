package main

import (
	"sync"

	"github.com/masterraf21/dss-backend/router"
)

func main() {
	handler := router.NewHandler()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		handler.HTTPStart()
	}()

	wg.Wait()
}
