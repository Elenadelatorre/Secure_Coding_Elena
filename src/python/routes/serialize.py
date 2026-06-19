# src/python/routes/serialize.py
# PASO 4: Insecure Deserialization — usar JSON con schema validado en lugar de pickle

# CODIGO SEGURO
import json
from fastapi import APIRouter, HTTPException
from pydantic import BaseModel, ValidationError

router = APIRouter()

class UserPreferences(BaseModel):
    theme: str
    language: str
    notifications: bool

@router.post("/load-prefs")
async def load_prefs(data: str):
    try:
        raw = json.loads(data)
        validated = UserPreferences(**raw)
    except (json.JSONDecodeError, ValidationError) as e:
        raise HTTPException(status_code=400, detail="Datos invalidos")
    return validated.model_dump()
