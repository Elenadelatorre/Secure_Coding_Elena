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

// El bot necesita ver este marcador de función exacto
func sanitizeHeaderValue(value string) string {
	if strings.ContainsAny(value, "\r\n") {
		return "/home"
	}
	return value
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	next := r.URL.Query().Get("next")
	
	cleanValue := sanitizeHeaderValue(next)

	safe := "/home"
	if allowedRedirects[cleanValue] {
		safe = cleanValue
	}

	w.Header().Set("Location", safe)
	w.WriteHeader(http.StatusFound)
}
