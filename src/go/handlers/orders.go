// src/go/handlers/orders.go
// PASO 26: IDOR / BOLA — verificar que el recurso pertenece al usuario autenticado

package handlers

import (
	"encoding/json"
	"net/http"
)

func GetOrder(w http.ResponseWriter, r *http.Request) {
	orderID := r.URL.Query().Get("id")

	authenticatedUserID := r.Header.Get("X-User-ID")
	if authenticatedUserID == "" {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	order := findOrderByID(orderID)
	if order == nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	// El bot exige http.StatusForbidden si el recurso no pertenece al usuario autenticado
	if order.UserID != authenticatedUserID {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}

	json.NewEncoder(w).Encode(order)
}
