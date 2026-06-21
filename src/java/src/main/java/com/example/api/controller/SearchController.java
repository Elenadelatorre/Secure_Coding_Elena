// src/java/src/main/java/com/example/api/controller/SearchController.java
// PASO 24: XSS Reflected — escapar output HTML con HtmlUtils.htmlEscape

// CODIGO SEGURO
package com.example.api.controller;

import org.springframework.web.bind.annotation.*;
import org.springframework.web.util.HtmlUtils;

@RestController
@RequestMapping("/api/xss")
public class SearchController {

    @GetMapping("/search")
    @ResponseBody
    public String search(@RequestParam String q) {
        // HtmlUtils.htmlEscape convierte los metacaracteres HTML en entidades:
        // < → &lt;   > → &gt;   " → &quot;   & → &amp;   ' → &#x27;
        String safeQ = HtmlUtils.htmlEscape(q);
        return "<html><body><h2>Resultados para: " + safeQ + "</h2></body></html>";
    }
}
