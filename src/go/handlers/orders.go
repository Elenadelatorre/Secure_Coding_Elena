// src/go/handlers/orders.go
// PASO 26: IDOR / BOLA — verificar que el recurso pertenece al usuario autenticado

// CODIGO SEGURO
package handlers

import (
    "encoding/json"
    "net/http"
)

func GetOrder(w http.ResponseWriter, r *http.Request) {
    orderID := r.URL.Query().Get("id")

    // El ID del usuario autenticado debe venir del middleware de autenticacion,
    // no de un parametro controlable por el cliente.
    // En produccion: extraer del JWT validado por el middleware.
    authenticatedUserID := r.Header.Get("X-User-ID")
    if authenticatedUserID == "" {
        http.Error(w, "unauthorized", http.StatusUnauthorized)
        return
    }

    order := findOrderByID(orderID)
    // Verificacion de propiedad: el pedido debe pertenecer al usuario autenticado
    if order == nil || order.UserID != authenticatedUserID {
        // Mismo error para "no existe" y "no es tuyo" — no revelar si el ID existe
        http.Error(w, "not found", http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(order)
}
