// src/go/handlers/middleware.go
// PASO 15: Clickjacking — middleware de cabeceras de seguridad con X-Frame-Options

// CODIGO SEGURO
package handlers

import "net/http"

func SecurityHeadersMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Prevenir clickjacking: rechazar iframe desde otros origenes
        w.Header().Set("X-Frame-Options", "DENY")

        // Alternativa moderna: CSP frame-ancestors (mas flexible)
        // w.Header().Set("Content-Security-Policy", "frame-ancestors 'none'")

        // Prevenir MIME sniffing: el navegador no adivina el Content-Type
        w.Header().Set("X-Content-Type-Options", "nosniff")

        // Limitar acceso a funcionalidades del navegador
        w.Header().Set("Permissions-Policy", "camera=(), microphone=(), geolocation=()")

        // Forzar HTTPS (HSTS)
        w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains")

        // Politica de referrer
        w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")

        next.ServeHTTP(w, r)
    })
}
