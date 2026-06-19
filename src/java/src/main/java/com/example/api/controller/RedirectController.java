// src/java/src/main/java/com/example/api/controller/RedirectController.java
// PASO 7: Open Redirect — allowlist de destinos de redireccion

package com.example.api.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;

// CODIGO SEGURO
@Controller
@RequestMapping("/auth")
public class RedirectController {

    private static final List<String> ALLOWED_REDIRECTS = List.of(
        "/dashboard",
        "/profile",
        "/settings",
        "/orders"
    );

    @GetMapping("/login")
    public String login(@RequestParam(defaultValue = "/dashboard") String next) {
        // Solo redirigir a rutas internas de la allowlist
        if (!ALLOWED_REDIRECTS.contains(next)) {
            return "redirect:/dashboard";  // destino seguro por defecto
        }
        return "redirect:" + next;
    }
}
