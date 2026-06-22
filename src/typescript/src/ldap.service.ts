// src/typescript/src/ldap.service.ts
// PASO 30: LDAP Injection — escapar metacaracteres de filtros LDAP

// CODIGO SEGURO
import { Injectable } from '@nestjs/common';

@Injectable()
export class LdapService {
  private readonly directory = [
    { uid: 'alice', userPassword: 'pass1', cn: 'Alice Smith', role: 'user' },
    { uid: 'admin', userPassword: 'adminpass', cn: 'Admin User', role: 'admin' },
  ];

  // Escapar los 5 metacaracteres especiales de filtros LDAP (RFC 4515)
  private escapeLdapFilter(value: string): string {
    return value
      .replace(/\\/g, '\\5C')   // \ → \5C  (debe ser primero para no escapar los otros escapes)
      .replace(/\*/g, '\\2A')   // * → \2A
      .replace(/\(/g, '\\28')   // ( → \28
      .replace(/\)/g, '\\29')   // ) → \29
      .replace(/\x00/g, '\\00'); // NUL → \00
  }

  private search(filter: string): Array<Record<string, string>> {
    const matchUid = filter.match(/\(uid=([^)]*)\)/);
    const matchPass = filter.match(/\(userPassword=([^)]*)\)/);
    if (!matchUid || !matchPass) return [];
    return this.directory.filter(
      e => e.uid === matchUid[1] && e.userPassword === matchPass[1],
    );
  }

  async authenticate(username: string, password: string): Promise<boolean> {
    // Escapar ANTES de interpoler en el filtro
    const safeUser = this.escapeLdapFilter(username);
    const safePass = this.escapeLdapFilter(password);
    const filter = `(&(uid=${safeUser})(userPassword=${safePass}))`;
    const results = this.search(filter);
    return results.length > 0;
  }
}
