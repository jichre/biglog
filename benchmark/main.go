package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

func main() {
	fmt.Println("start")

	globalWaiter := new(sync.WaitGroup)

	starTimestamp := time.Now()

	for n := 1; n <= 5; n++ {
		globalWaiter.Add(1)
		go func(num int) {
			fmt.Printf("start goruntine num=%v\n", num)
			for i := 0; i < 100; i++ {
				resp, err := http.Post("http://127.0.0.1:80/api/postLog",
					"application/x-www-form-urlencoded",
					strings.NewReader("name=cjb"))
				if err != nil {
					fmt.Println(err)
				}

				defer resp.Body.Close()
			}
			globalWaiter.Done()
			fmt.Printf("done goruntine num=%v\n", num)
		}(n)
	}

	globalWaiter.Wait()
	endTimestamp := time.Now()
	timeCost := endTimestamp.Sub(starTimestamp)
	fmt.Printf("Cost: %v\n", timeCost)
	fmt.Println("Complete")
}
