// src/typescript/src/proxy.controller.ts
// PASO 28: SSRF — validar host contra allowlist antes de proxificar requests

// CODIGO SEGURO
import { BadRequestException, Controller, Get, Query } from '@nestjs/common';
import axios from 'axios';

// Lista blanca de hostnames permitidos
const ALLOWED_HOSTS = new Set([
  'api.github.com',
  'jsonplaceholder.typicode.com',
  'api.openweathermap.org',
]);

@Controller('proxy')
export class ProxyController {
  @Get('/fetch')
  async fetch(@Query('url') url: string): Promise<string> {
    if (!url) {
      throw new BadRequestException('URL requerida');
    }

    // Parsear y validar la URL antes de hacer la peticion
    let parsed: URL;
    try {
      parsed = new URL(url);
    } catch {
      throw new BadRequestException('URL malformada');
    }

    // Solo HTTPS — previene file://, ftp://, gopher://, etc.
    if (parsed.protocol !== 'https:') {
      throw new BadRequestException('Solo se permite HTTPS');
    }

    // Allowlist de hostnames — rechaza IMDS, red interna, localhost
    if (!ALLOWED_HOSTS.has(parsed.hostname)) {
      throw new BadRequestException('Host no permitido');
    }

    // maxRedirects: 0 — no seguir redirects que puedan apuntar a hosts internos
    const response = await axios.get(url, { maxRedirects: 0 });
    return response.data;
  }
}
