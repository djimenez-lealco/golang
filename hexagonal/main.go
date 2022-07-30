package main

import (
	"hexagonal/infraestructure/entrypoints/api"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go api.Run(&wg)
	wg.Wait()
}
