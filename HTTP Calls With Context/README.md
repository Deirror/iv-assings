# Concurrent HTTP Calls with Timeout & Context Cancellation

Scenario
-

You’re writing a Go service that queries multiple external HTTP endpoints in parallel.

But:
- If even one call fails or times out,
- cancel all others and report an error.

Task Requirements
-

- Given a slice of URLs, launch one goroutine per HTTP GET request.
- Use context.WithTimeout or context.WithCancel to cancel all if any single request fails or takes too long.
- Collect successful responses into a slice.
- If all succeed, print the combined results; if any fail, print an error.