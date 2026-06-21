// src/java/src/main/java/com/example/api/controller/CommentsController.java
// PASO 25: XSS Stored — escapar contenido almacenado al renderizarlo en HTML

// CODIGO SEGURO
package com.example.api.controller;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.util.HtmlUtils;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;

@RestController
@RequestMapping("/api/comments")
public class CommentsController {

    private final List<String> comments = new ArrayList<>();

    @PostMapping
    public ResponseEntity<?> addComment(@RequestBody Map<String, String> body) {
        String comment = body.get("comment");
        if (comment == null || comment.isBlank()) {
            return ResponseEntity.badRequest().body("Comentario vacio");
        }
        comments.add(comment);
        return ResponseEntity.ok().build();
    }

    @GetMapping
    @ResponseBody
    public String getComments() {
        StringBuilder sb = new StringBuilder("<ul>");
        for (String c : comments) {
            // Escapar en el punto de salida: convertir < > " & ' en entidades HTML
            sb.append("<li>").append(HtmlUtils.htmlEscape(c)).append("</li>");
        }
        sb.append("</ul>");
        return sb.toString();
    }
}
