package main

import (
	"fmt"
	"sync"

	"github.com/masterraf21/dss-backend/router"
)

func main() {
	// ctx := context.TODO()
	handler := router.NewHandler()
	// handler.HTTPStart()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		fmt.Printf("RUNNING SERVER")
		defer wg.Done()
		handler.HTTPStart()
	}()

	wg.Wait()
}
