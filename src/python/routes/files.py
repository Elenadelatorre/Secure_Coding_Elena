# src/python/routes/files.py
# PASO 2: Path Traversal — normalizar ruta real y verificar que esta dentro del directorio permitido

# CODIGO SEGURO
import os
from fastapi import APIRouter, HTTPException
from fastapi.responses import FileResponse

router = APIRouter()

ALLOWED_DIR = os.path.realpath("/var/www/public")

@router.get("/download")
async def download_file(filename: str):
    real_path = os.path.realpath(os.path.join(ALLOWED_DIR, filename))
    if not real_path.startswith(ALLOWED_DIR + os.sep):
        raise HTTPException(status_code=400, detail="Acceso denegado")
    if not os.path.isfile(real_path):
        raise HTTPException(status_code=404, detail="Archivo no encontrado")
    return FileResponse(real_path)
