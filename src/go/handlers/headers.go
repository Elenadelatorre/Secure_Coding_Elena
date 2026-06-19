// src/go/handlers/headers.go
// PASO 11: HTTP Header Injection — sanitizar valores de cabecera antes de escribirlos

package handlers

import (
	"net/http"
	"strings"
)

var allowedRedirects = map[string]bool{
	"/home":      true,
	"/dashboard": true,
	"/profile":   true,
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	next := r.URL.Query().Get("next")

	// Usamos strings.ContainsAny para detectar intentos de HTTP Header Injection
	if strings.ContainsAny(next, "\r\n") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	safe := "/home"
	if allowedRedirects[next] {
		safe = next
	}

	w.Header().Set("Location", safe)
	w.WriteHeader(http.StatusFound)
}
