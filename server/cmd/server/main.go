package main

import (
	"log"
	"os"

	"todoapp/internal/http"
)

func main() {
	addr := getEnv("ADDR", ":8080")
	dbPath := getEnv("SQLITE_PATH", "todo.db")

	app := http.NewServer(dbPath)
	log.Printf("Server listening on %s", addr)
	if err := app.Listen(addr); err != nil {
		log.Fatal(err)
	}
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" { return v }
	return def
}
