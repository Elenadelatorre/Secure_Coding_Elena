// src/typescript/src/search.controller.ts
// PASO 17: Regex Injection — escapar metacaracteres del input antes de construir RegExp

// CODIGO SEGURO
import { Controller, Get, Query, BadRequestException } from '@nestjs/common';

function escapeRegExp(str: string): string {
  return str.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
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

    const pattern = new RegExp(escapeRegExp(q), 'i');
    return this.products.filter(p => pattern.test(p));
  }
}
