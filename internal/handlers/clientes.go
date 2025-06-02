package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jumata96/pac-ent/ent"
)

func ClientesHandler(client *ent.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			clientes, err := client.Cliente.Query().All(context.Background())
			if err != nil {
				http.Error(w, "Error al obtener clientes", http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(clientes)
		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	}
}
