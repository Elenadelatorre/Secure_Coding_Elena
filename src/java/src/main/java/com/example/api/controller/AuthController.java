// src/java/src/main/java/com/example/api/controller/AuthController.java
// PASO 9: Log Injection — sanitizar input antes de escribirlo en logs

package com.example.api.controller;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.Map;

@RestController
@RequestMapping("/api/auth")
public class AuthController {

    // CODIGO SEGURO
private static String sanitizeForLog(String input) {
    if (input == null) return "null";
    // Eliminar caracteres de control (saltos de linea, tabuladores, etc.)
    String sanitized = input.replaceAll("[\\r\\n\\t]", "_");
    // Limitar longitud para evitar logs anormalmente largos
    if (sanitized.length() > 100) {
        sanitized = sanitized.substring(0, 100) + "[truncado]";
    }
    return sanitized;
}

@PostMapping("/login")
public ResponseEntity<?> login(@RequestParam String username,
                               @RequestParam String password) {
    // Usar logging parametrizado Y sanitizacion del input
    log.info("Login attempt for user: {}", sanitizeForLog(username));
    return ResponseEntity.ok(Map.of("message", "OK"));
}
