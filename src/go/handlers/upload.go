// src/go/handlers/upload.go
// PASO 12: Race Condition (TOCTOU) — operacion atomica con O_EXCL elimina la ventana

package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const uploadDir = "/var/uploads"

// CODIGO SEGURO
package handlers

import (
    "io"
    "net/http"
    "os"
    "path/filepath"
)

const uploadDir = "/var/uploads"

func UploadHandler(w http.ResponseWriter, r *http.Request) {
    // Usar solo el nombre base del archivo (elimina path traversal)
    filename := filepath.Base(r.FormValue("name"))
    if filename == "." || filename == "" {
        http.Error(w, "Invalid filename", http.StatusBadRequest)
        return
    }
    path := filepath.Join(uploadDir, filename)

    // O_CREATE|O_EXCL es atomico: crea el archivo SOLO si no existe.
    // Si ya existe, falla con error. No hay ventana entre check y create.
    f, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0644)
    if err != nil {
        if os.IsExist(err) {
            http.Error(w, "File already exists", http.StatusConflict)
        } else {
            http.Error(w, "Internal error", http.StatusInternalServerError)
        }
        return
    }
    defer f.Close()
    io.Copy(f, r.Body)
    w.WriteHeader(http.StatusCreated)
}
