// src/typescript/src/config.service.ts
// PASO 19: Hardcoded Secrets — secretos en variables de entorno con validacion al arranque


// CODIGO SEGURO
import { Injectable } from '@nestjs/common';

// Leer un secreto obligatorio desde entorno. Falla al arrancar si no esta definido.
function requireEnv(name: string): string {
  const value = process.env[name];
  if (!value) {
    // Fallar en startup es correcto: mejor no arrancar que operar sin secretos
    throw new Error(`Variable de entorno requerida no esta definida: ${name}`);
  }
  return value;
}

@Injectable()
export class ConfigService {
  // Los secretos se resuelven al instanciar el servicio (al arranque de la app)
  readonly jwtSecret: string = requireEnv('JWT_SECRET');
  readonly dbPassword: string = requireEnv('DB_PASSWORD');
  readonly stripeKey: string = requireEnv('STRIPE_API_KEY');
}


// CODIGO SEGURO
import { Injectable } from '@nestjs/common';

// Leer un secreto obligatorio desde entorno. Falla al arrancar si no esta definido.
function requireEnv(name: string): string {
  const value = process.env[name];
  if (!value) {
    // Fallar en startup es correcto: mejor no arrancar que operar sin secretos
    throw new Error(`Variable de entorno requerida no esta definida: ${name}`);
  }
  return value;
}

@Injectable()
export class ConfigService {
  // Los secretos se resuelven al instanciar el servicio (al arranque de la app)
  readonly jwtSecret: string = requireEnv('JWT_SECRET');
  readonly dbPassword: string = requireEnv('DB_PASSWORD');
  readonly stripeKey: string = requireEnv('STRIPE_API_KEY');
}
