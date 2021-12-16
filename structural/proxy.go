package main

import "fmt"

type server interface {
	handleRequest(string, string) (int, string)
}

type nginx struct {
	application        *application
	maxAllowedRequests int
	rateLimiter        map[string]int
}

func newNginxServer() *nginx {
	return &nginx{
		application:        &application{},
		maxAllowedRequests: 2,
		rateLimiter:        make(map[string]int, 1),
	}
}

func (n *nginx) handleRequest(path, method string) (int, string) {
	allowed := n.checkRateLimit(path)
	if !allowed {
		return 403, "Not Allowed, Rate limited"
	}
	return n.application.handleRequest(path, method)
}

func (n *nginx) checkRateLimit(path string) bool {
	n.rateLimiter[path]++
	if n.rateLimiter[path] > n.maxAllowedRequests {
		return false
	}
	return true
}

type application struct{}

func (a *application) handleRequest(path, method string) (int, string) {
	if path == "/api/status" && method == "GET" {
		return 200, "OK"
	}
	return 404, "Not Found"
}

// CLIENT
func main() {
	nginxServer := newNginxServer()
	httpCode, body := nginxServer.handleRequest("/api/status", "GET")
	fmt.Printf("\nHttpCode: %d\nBody: %s\n", httpCode, body)
	httpCode, body = nginxServer.handleRequest("/api/status", "GET")
	fmt.Printf("\nHttpCode: %d\nBody: %s\n", httpCode, body)
	httpCode, body = nginxServer.handleRequest("/api/status", "GET")
	fmt.Printf("\nHttpCode: %d\nBody: %s\n", httpCode, body)

	httpCode, body = nginxServer.handleRequest("/nope", "POST")
	fmt.Printf("\nHttpCode: %d\nBody: %s\n", httpCode, body)
}
