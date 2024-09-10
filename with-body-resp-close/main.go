package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"sync/atomic"
)

var totalRequests uint64

func requestWithBodyClose(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		resp, err := http.Get("https://www.baidu.com")
		if err != nil {
			log.Fatalf("HTTP request error: %v", err)
		}
		_ = resp // without close

		requests := atomic.AddUint64(&totalRequests, 1)
		fmt.Printf("Request successful, total requests: %d\n", requests)
	}
}

func main() {
	go func() {
		log.Fatal(http.ListenAndServe(":8081", nil))
	}()

	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go requestWithBodyClose(&wg)
	}

	wg.Wait()
}
