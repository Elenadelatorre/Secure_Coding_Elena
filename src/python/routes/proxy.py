# src/python/routes/proxy.py
# PASO 23: SSRF — validar host contra allowlist antes de hacer requests externos

# CODIGO SEGURO
import requests
from urllib.parse import urlparse
from fastapi import APIRouter, HTTPException

router = APIRouter()

# Lista blanca de hosts permitidos — solo IPs/dominios explicitos
ALLOWED_HOSTS = {
    "api.github.com",
    "api.openweathermap.org",
    "jsonplaceholder.typicode.com",
}

def _validate_ssrf(url: str) -> None:
    """Valida que la URL sea segura antes de hacer la peticion."""
    try:
        parsed = urlparse(url)
    except Exception:
        raise HTTPException(status_code=400, detail="URL malformada")

    # Solo HTTPS: previene protocolos como file://, ftp://, gopher://
    if parsed.scheme not in ("https",):
        raise HTTPException(status_code=400, detail="Solo se permite HTTPS")

    # Allowlist de hosts: rechaza IMDS, red interna, localhost
    if parsed.hostname not in ALLOWED_HOSTS:
        raise HTTPException(status_code=400, detail="Host no permitido")


@router.get("/fetch")
async def fetch_url(url: str):
    _validate_ssrf(url)
    # allow_redirects=False: previene bypasses via redirects a hosts internos
    response = requests.get(url, timeout=5, allow_redirects=False)
    return {"status": response.status_code, "content": response.text[:500]}
