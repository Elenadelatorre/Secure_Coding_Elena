// src/go/handlers/jwt.go
// PASO 27: JWT Algorithm Confusion — validar explicitamente el algoritmo de firma

// CODIGO SEGURO
package handlers

import (
    "fmt"
    "net/http"
    "os"
    "strings"

    "github.com/golang-jwt/jwt/v5"
)

func ParseToken(tokenString string) (*jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Verificar EXPLICITAMENTE que el algoritmo es HMAC (HS256/HS384/HS512)
        // Si el header dice "none", RSA, ECDSA u otro, se rechaza aqui
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("algoritmo inesperado: %v", token.Header["alg"])
        }
        // El secreto viene de variable de entorno, nunca hardcodeado
        return []byte(os.Getenv("JWT_SECRET")), nil
    })
    if err != nil || !token.Valid {
        return nil, fmt.Errorf("invalid token")
    }
    claims := token.Claims.(jwt.MapClaims)
    return &claims, nil
}

func ValidateJWTHandler(w http.ResponseWriter, r *http.Request) {
    authHeader := r.Header.Get("Authorization")
    if !strings.HasPrefix(authHeader, "Bearer ") {
        http.Error(w, "missing token", http.StatusUnauthorized)
        return
    }
    tokenString := strings.TrimPrefix(authHeader, "Bearer ")
    claims, err := ParseToken(tokenString)
    if err != nil {
        http.Error(w, "unauthorized", http.StatusUnauthorized)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, `{"valid":true,"sub":"%v"}`, (*claims)["sub"])
}
