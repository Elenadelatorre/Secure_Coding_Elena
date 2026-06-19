# src/python/routes/cors.py
# PASO 5: CORS misconfiguration — origen especifico desde variable de entorno,
#         sin wildcard cuando allow_credentials=True

# CODIGO SEGURO
import os
from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware

ALLOWED_ORIGINS = [
    origin.strip()
    for origin in os.environ.get("CORS_ALLOWED_ORIGINS", "https://app.empresa.com").split(",")
    if origin.strip()
]

def configure_cors(app: FastAPI) -> None:
    app.add_middleware(
        CORSMiddleware,
        allow_origins=ALLOWED_ORIGINS,             # lista explicita, no wildcard
        allow_credentials=True,
        allow_methods=["GET", "POST", "PUT", "DELETE"],
        allow_headers=["Authorization", "Content-Type"],
    )
