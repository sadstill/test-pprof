package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func requestWithNoBodyClose() error {
	_, err := http.Get("https://www.baidu.com")
	if err != nil {
		return fmt.Errorf("http.Get failed: %w", err)
	}

	return nil
}

func main() {
	go func() {
		log.Fatal(http.ListenAndServe(":8082", nil))
	}()

	step := 0

	for {
		time.Sleep(time.Microsecond * 100)

		step++

		err := requestWithNoBodyClose()
		if err != nil {
			fmt.Printf("[%d] requestNoClose failed: %s", step, err)
			continue
		}

		fmt.Printf("[%d] ok\n", step)
	}
}
