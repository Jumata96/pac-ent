package main

import (
	"context"
	"fmt"
	"log"

	// "github.com/tuusuario/pac-ent/ent" // Asegúrate de que esto coincida con tu módulo
	"pac-ent/ent"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent.db?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	user, err := client.User.Create().SetName("Julinho").SetEmail("julinho@example.com").Save(ctx)
	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}
	fmt.Println("User:", user)
}
