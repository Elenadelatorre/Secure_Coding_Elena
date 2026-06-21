// src/typescript/src/merge.controller.ts
// PASO 16: Prototype Pollution — validar con DTO y Object.create(null) en lugar de Object.assign

// CODIGO SEGURO
import { Body, Controller, Post } from '@nestjs/common';
import { plainToInstance } from 'class-transformer';
import { IsBoolean, IsOptional, IsString, validateSync } from 'class-validator';

class UserPreferencesDto {
  @IsOptional()
  @IsString()
  theme?: string;

  @IsOptional()
  @IsString()
  language?: string;

  @IsOptional()
  @IsBoolean()
  notifications?: boolean;
}

@Controller('user')
export class MergeController {
  @Post('/preferences')
  updatePreferences(@Body() body: Record<string, unknown>): Record<string, unknown> {
    // Validar contra DTO tipado: rechaza claves no definidas en la clase
    const dto = plainToInstance(UserPreferencesDto, body);
    const errors = validateSync(dto);
    if (errors.length > 0) {
      throw new Error('Invalid input');
    }

    // Object.create(null) crea un objeto SIN prototipo (no hereda de Object.prototype)
    // Un atacante no puede contaminar Object.prototype a traves de este objeto
    const prefs = Object.create(null) as Record<string, unknown>;
    if (dto.theme !== undefined) prefs['theme'] = dto.theme;
    if (dto.language !== undefined) prefs['language'] = dto.language;
    if (dto.notifications !== undefined) prefs['notifications'] = dto.notifications;
    return prefs;
  }
}
