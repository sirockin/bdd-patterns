package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/sirockin/cucumber-screenplay-go/features/driver/domain"
	"github.com/sirockin/cucumber-screenplay-go/internal/server"
)

func main() {
	var port = flag.Int("port", 8080, "port to run server on")
	flag.Parse()

	// Create domain application
	app := domain.New()

	// Create HTTP server wrapping the domain
	httpServer := server.NewServer(app)

	// Start server
	addr := fmt.Sprintf(":%d", *port)
	log.Printf("Starting server on http://localhost%s", addr)
	log.Printf("API endpoints:")
	log.Printf("  POST   /accounts")
	log.Printf("  GET    /accounts/{name}")
	log.Printf("  POST   /accounts/{name}/activate")
	log.Printf("  POST   /accounts/{name}/authenticate")
	log.Printf("  GET    /accounts/{name}/authentication-status")
	log.Printf("  GET    /accounts/{name}/projects")
	log.Printf("  POST   /accounts/{name}/projects")
	log.Printf("  DELETE /clear")

	if err := http.ListenAndServe(addr, httpServer); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}