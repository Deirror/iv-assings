# Concurrent Squaring with Goroutine per Task + WaitGroup

Description
-

Instead of using fixed worker goroutines,
- launch one goroutine per number
- use a sync.WaitGroup to wait for all goroutines to finish
- collect squared results in an ordered slice

Requirements
-

- Use sync.WaitGroup to wait for all goroutines
- Results must print in the same order as input
- No worker pool, just spin up one goroutine per item