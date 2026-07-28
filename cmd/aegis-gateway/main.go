package main

import (
	"log"
	"net/http"

	"github.com/hansunwoo232/Aegis-Vari/internal/smt"
)

func main() {
	log.Println("Initializing Aegis-Vari ePCA Engine...")

	// Initialize the Z3 SMT Solver
	solver, err := smt.NewZ3Solver()
	if err != nil {
		log.Fatalf("Failed to start SMT solver: %v", err)
	}
	defer solver.Close()

	log.Println("ePCA Engine active. First-order logic verification is ready.")

	// Placeholder for the MCP Security Gateway
	log.Println("MCP Security Gateway listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Gateway crashed: %v", err)
	}
}
