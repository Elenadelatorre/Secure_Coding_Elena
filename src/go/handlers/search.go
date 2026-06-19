// src/go/handlers/search.go
// PASO 13: ReDoS — regex seguro y longitud maxima de input

// CODIGO SEGURO
package handlers

import (
    "net/http"
    "regexp"
)

const maxInputLength = 254  // longitud maxima de email segun RFC 5321

// Patron lineal: sin cuantificadores anidados, sin alternaciones con prefijo comun
// [a-zA-Z0-9._%+\-]+ matchea una vez por caracter, sin posibilidad de repartir
var safeEmailPattern = regexp.MustCompile(
    `^[a-zA-Z0-9._%+\-]+@example\.com$`,
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
    input := r.URL.Query().Get("q")

    // Limitar longitud antes de evaluar la regex
    if len(input) > maxInputLength {
        http.Error(w, "Input too long", http.StatusBadRequest)
        return
    }

    if safeEmailPattern.MatchString(input) {
        w.Write([]byte("valid"))
    } else {
        w.Write([]byte("invalid"))
    }
}
