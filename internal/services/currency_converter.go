package services

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/shopspring/decimal"
)

type Result struct {
	Pair string
	Rate decimal.Decimal
	Err  error
}

var fixedRates = map[string]decimal.Decimal{
	"USD/KZT": decimal.NewFromFloat(540.42),
	"RUB/KZT": decimal.NewFromFloat(6.71),
	"USD/RUB": decimal.NewFromFloat(80.26),
	"EUR/KZT": decimal.NewFromFloat(590.10),
	"EUR/USD": decimal.NewFromFloat(1.15),
}

func fetchRate(ctx context.Context, pair string) Result {
	ch := make(chan Result, 1)
	go func() {
		delay := time.Duration(rand.Intn(2000)) * time.Millisecond
		time.Sleep(delay)
		rate, ok := fixedRates[pair]
		if !ok {
			ch <- Result{Pair: pair, Err: fmt.Errorf("unknown pair")}
			return
		}
		ch <- Result{Pair: pair, Rate: rate, Err: nil}
	}()

	select {
	case res := <-ch:
		return res
	case <-ctx.Done():
		return Result{Pair: pair, Err: fmt.Errorf("запрос превысил таймаут")}
	}
}

func RunCurrencyConverter() {
	pairs := []string{"USD/KZT", "RUB/KZT", "USD/RUB", "EUR/KZT", "EUR/USD"}
	results := make(chan Result, len(pairs))
	var wg sync.WaitGroup
	store := make(map[string]decimal.Decimal)
	var mu sync.Mutex

	for _, pair := range pairs {
		wg.Add(1)
		go func(pair string) {
			defer wg.Done()
			ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
			defer cancel()
			res := fetchRate(ctx, pair)
			results <- res
		}(pair)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Println("Полученные курсы валют:")
	for res := range results {
		if res.Err == nil {
			mu.Lock()
			store[res.Pair] = res.Rate
			mu.Unlock()
			fmt.Printf("%s: %s\n", res.Pair, res.Rate.String())
		} else {
			fmt.Printf("%s: %s\n", res.Pair, res.Err.Error())
		}
	}
}

func RunCurrencyConverterAPI() map[string]interface{} {
	pairs := []string{"USD/KZT", "RUB/KZT", "USD/RUB", "EUR/KZT", "EUR/USD"}
	results := make(chan Result, len(pairs))
	var wg sync.WaitGroup
	store := make(map[string]interface{})
	var mu sync.Mutex

	for _, pair := range pairs {
		wg.Add(1)
		go func(pair string) {
			defer wg.Done()
			ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
			defer cancel()
			res := fetchRate(ctx, pair)
			results <- res
		}(pair)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		if res.Err == nil {
			mu.Lock()
			store[res.Pair] = res.Rate.String()
			mu.Unlock()
		} else {
			mu.Lock()
			store[res.Pair] = res.Err.Error()
			mu.Unlock()
		}
	}
	return store
}
