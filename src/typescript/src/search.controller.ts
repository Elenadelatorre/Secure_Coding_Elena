// src/typescript/src/search.controller.ts
// PASO 17: Regex Injection — escapar metacaracteres del input antes de construir RegExp

// CODIGO SEGURO
import { Controller, Get, Query, BadRequestException } from '@nestjs/common';

// Escapar todos los metacaracteres de regex en el input del usuario
function escapeRegExp(str: string): string {
  return str.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
  // Los caracteres . * + ? ^ $ { } ( ) | [ ] \ son escapados con \
  // Asi, '.' se convierte en '\.' que solo matchea un punto literal
}

@Controller('products')
export class SearchController {
  private readonly products = ['laptop', 'phone', 'tablet', 'monitor', 'keyboard'];

  @Get('/search')
  search(@Query('q') q: string): string[] {
    if (!q || q.length === 0) {
      return [];
    }
    if (q.length > 100) {
      throw new BadRequestException('Query demasiado larga');
    }

    // Escapar metacaracteres: el usuario solo puede buscar texto literal
    const safePattern = escapeRegExp(q);
    const pattern = new RegExp(safePattern, 'i');
    return this.products.filter(p => pattern.test(p));
  }
}
