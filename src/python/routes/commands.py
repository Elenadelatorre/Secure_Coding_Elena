# src/python/routes/commands.py
# PASO 1: Command Injection — subprocess sin shell=True y validacion de hostname

# ✅ CÓDIGO SEGURO
import re
import subprocess
from fastapi import APIRouter, HTTPException

router = APIRouter()

# Allowlist estricta: solo hostnames/IPs con caracteres legítimos
VALID_HOSTNAME = re.compile(r'^[a-zA-Z0-9.\-]{1,253}$')

@router.get("/ping")
async def ping_host(host: str):
    if not VALID_HOSTNAME.match(host):
        raise HTTPException(status_code=400, detail="Hostname inválido")
    result = subprocess.run(
        ["ping", "-c", "1", host],  # ✅ lista de args → execvp() directo, sin shell
        capture_output=True,
        text=True,
        timeout=5
    )
    return {"output": result.stdout}
