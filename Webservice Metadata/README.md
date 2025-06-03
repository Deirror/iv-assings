# Cloud Microservice with Go

Task Description
-

You are to implement a simple Go microservice that:
- Exposes a REST API endpoint
- On request, fetches some mocked customer metadata
- Returns it as a JSON response
- Logs each request (with timestamp) to the console
- Supports concurrent requests safely

Requirements:
- Use Go (Golang)
- Use Go’s built-in HTTP package to set up the server
- Store mocked customer metadata in-memory (e.g., a map[int]string for customer IDs - names)
- Use a mutex or atomic.Value to make sure the map is safe for concurrent reads (even though it’s mocked, show you understand concurrency)
- When a request comes to /customer/{id}, return the customer’s name in JSON { "id": X, "name": "..." }
- If the customer ID is not found, return HTTP 404
- Log each request’s ID and timestamp to the console

Bonus (if you want to stretch a bit):
- Add a /customers endpoint that lists all customers
- Wrap the server in a Docker container (provide a Dockerfile)