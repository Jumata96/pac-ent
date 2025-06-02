package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/jumata96/pac-ent/ent"
	"github.com/jumata96/pac-ent/internal/handlers"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando .env")
	}

	// Formar DSN de PostgreSQL
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("falló conexión a postgres: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("falló la creación de esquema: %v", err)
	}

	http.HandleFunc("/clientes", handlers.ClientesHandler(client))
	log.Println("Servidor escuchando en :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
