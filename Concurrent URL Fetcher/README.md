# Concurrent URL Fetcher

Description
-

You are given a list of URLs.

Your task is to write a Go program that:
 - Spawns multiple worker goroutines (limit to, say, 5 concurrent)
 - Each worker fetches the content of a URL using HTTP GET
 - Collects the HTTP status code (e.g., 200, 404) for each URL
 - Reports the status codes in the same order as the original URL list

Constraints:

 - Use Go channels or sync tools.
 - Respect the concurrency limit (no more than 5 at the same time).
 - Ensure results are ordered, even though fetches happen in parallel.