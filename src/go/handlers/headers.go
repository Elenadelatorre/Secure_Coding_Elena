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

// sanitizeHeaderValue elimina caracteres de control del valor de un header
func sanitizeHeaderValue(value string) string {
    value = strings.ReplaceAll(value, "\r", "")
    value = strings.ReplaceAll(value, "\n", "")
    return value
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
    next := r.URL.Query().Get("next")
    sanitized := sanitizeHeaderValue(next)

    // Allowlist: solo redirigir a rutas internas conocidas
    if !allowedRedirects[sanitized] {
        sanitized = "/home"
    }

    w.Header().Set("Location", sanitized)
    w.WriteHeader(http.StatusFound)
}
