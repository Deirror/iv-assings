package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	urls := []string{
		"https://httpbin.org/delay/1",
		"https://httpbin.org/delay/2",
		"https://httpbin.org/delay/3",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	wg := &sync.WaitGroup{}

	ch := make(chan string, len(urls))
	errCh := make(chan error, 1) // only need 1 slot, we cancel on first error

	go func() {
		select {
		case err := <-errCh:
			fmt.Println("Error:", err)
			cancel() // cancel the context for all
		case <-ctx.Done():
			// context timed out, nothing to do here
		}
	}()

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
			if err != nil {
				select {
				case errCh <- err:
				default:
				}
				return
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				select {
				case errCh <- err:
				default:
				}
				return
			}
			defer resp.Body.Close()

			if ctx.Err() != nil {
				return
			}

			select {
			case ch <- resp.Status:
			case <-ctx.Done():
				return
			}
		}(url)
	}

	wg.Wait()
	close(ch)

	for res := range ch {
		fmt.Println("Response status:", res)
	}
}
