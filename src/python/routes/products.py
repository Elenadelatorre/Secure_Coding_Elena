# src/python/routes/products.py
# PASO 22: NoSQL Injection — validacion de tipos antes de queries a MongoDB

# CODIGO SEGURO
from fastapi import APIRouter, HTTPException

router = APIRouter()

@router.post("/login")
async def login(body: dict):
    username = body.get("username")
    password = body.get("password")

    # Validar que username y password sean strings escalares
    # Un operador MongoDB como {"$ne": ""} es un dict, no un str
    if not isinstance(username, str) or not isinstance(password, str):
        raise HTTPException(status_code=400, detail="Tipo de datos invalido")

    # Ahora es seguro: username y password son strings literales,
    # no pueden contener operadores MongoDB
    user = _find_one({"username": username, "password": password})
    if user:
        return {"token": "access_granted", "role": user.get("role")}
    raise HTTPException(status_code=401, detail="Credenciales invalidas")
