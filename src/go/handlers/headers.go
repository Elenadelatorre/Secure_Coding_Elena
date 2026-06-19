// src/go/handlers/headers.go
// PASO 11: HTTP Header Injection — sanitizar valores de cabecera antes de escribirlos

// CODIGO SEGURO
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

func sanitizeHeaderValue(value string) string {
	value = strings.ReplaceAll(value, "\r", "")
	value = strings.ReplaceAll(value, "\n", "")
	return value
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	next := r.URL.Query().Get("next")
	sanitized := sanitizeHeaderValue(next)

	// Por defecto usamos la ruta segura
	safe := "/home"

	// Si está en la lista permitida, la usamos
	if allowedRedirects[sanitized] {
		safe = sanitized
	}

	w.Header().Set("Location", safe)
	w.WriteHeader(http.StatusFound)
}
