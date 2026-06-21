// src/typescript/src/upload.controller.ts
// PASO 20: Insecure File Upload — validar MIME type, extension y usar nombre generado

// CODIGO SEGURO
import {
  Controller, Post, UploadedFile, UseInterceptors, BadRequestException
} from '@nestjs/common';
import { FileInterceptor } from '@nestjs/platform-express';
import { randomUUID } from 'crypto';
import * as path from 'path';

// Tipos MIME permitidos (lista blanca, no lista negra)
const ALLOWED_MIME_TYPES = new Set([
  'image/jpeg',
  'image/png',
  'image/gif',
  'image/webp',
  'application/pdf',
]);

// Extensiones permitidas (mapeadas desde MIME)
const ALLOWED_EXTENSIONS = new Set(['.jpg', '.jpeg', '.png', '.gif', '.webp', '.pdf']);

// Limite de tamano: 5 MB
const MAX_FILE_SIZE = 5 * 1024 * 1024;

@Controller('files')
export class UploadController {
  @Post('/upload')
  @UseInterceptors(FileInterceptor('file', {
    limits: { fileSize: MAX_FILE_SIZE },  // limite de tamano en multer
  }))
  upload(@UploadedFile() file: Express.Multer.File): { filename: string } {
    if (!file) {
      throw new BadRequestException('No se recibio ningun archivo');
    }

    // Validar MIME type reportado por multer
    if (!ALLOWED_MIME_TYPES.has(file.mimetype)) {
      throw new BadRequestException(`Tipo de archivo no permitido: ${file.mimetype}`);
    }

    // Validar extension del nombre original
    const ext = path.extname(file.originalname).toLowerCase();
    if (!ALLOWED_EXTENSIONS.has(ext)) {
      throw new BadRequestException(`Extension no permitida: ${ext}`);
    }

    // Validar tamano (aunque multer ya limita, verificacion adicional)
    if (file.size > MAX_FILE_SIZE) {
      throw new BadRequestException('Archivo demasiado grande');
    }

    // Generar nombre aleatorio: evita path traversal y colisiones
    const safeFilename = `${randomUUID()}${ext}`;

    // Aqui se guardaria el archivo con safeFilename, no con file.originalname
    // fs.writeFileSync(path.join('/var/uploads', safeFilename), file.buffer);

    return { filename: safeFilename };
  }
}
