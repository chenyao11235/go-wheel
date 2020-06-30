package main

import (
	"context"
	"log"
	"sync"

	"golang.org/x/time/rate"
)

// 只能精确到秒

type APIConnection struct {
	rateLimiter *rate.Limiter
}

func (a *APIConnection) ReadFile(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil {
		return err
	}
	return nil
}

func (a *APIConnection) ResolveAddress(ctx context.Context) error {
	if err := a.rateLimiter.Wait(ctx); err != nil {
		return err
	}
	return nil
}

func Open() *APIConnection {
	return &APIConnection{
		// 最高每秒钟10个
		rateLimiter: rate.NewLimiter(rate.Limit(10), 5),
	}
}

func main() {
	apiConnection := Open()
	var wg = sync.WaitGroup{}
	wg.Add(20)

	for i := 0; i < 20; i++ {
		go func() {
			defer wg.Done()
			err := apiConnection.ReadFile(context.Background())
			if err != nil {
				log.Printf("Cannot readfile: %v", err)
			}
			log.Printf("ReadFile")
		}()
	}

	for i := 0; i < 20; i++ {
		go func() {
			defer wg.Done()
			err := apiConnection.ResolveAddress(context.Background())
			if err != nil {
				log.Printf("Cannot ResolveAddress: %v", err)
			}
			log.Printf("ResolveAddress")
		}()
	}
	wg.Wait()
}
