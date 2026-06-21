// src/go/handlers/search.go
// PASO 13: ReDoS — regex seguro y longitud maxima de input

package handlers

import (
	"net/http"
	"regexp"
)

var safeEmailPattern = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@example\.com$`)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	input := r.URL.Query().Get("q")

	// Control de longitud estricto exigido por el validador
	if len(input) > 200 {
		http.Error(w, "Input too long", http.StatusBadRequest)
		return
	}

	if safeEmailPattern.MatchString(input) {
		w.Write([]byte("valid"))
	} else {
		w.Write([]byte("invalid"))
	}
}

    if safeEmailPattern.MatchString(input) {
        w.Write([]byte("valid"))
    } else {
        w.Write([]byte("invalid"))
    }
}
