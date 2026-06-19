import json
from fastapi import APIRouter, HTTPException
from pydantic import BaseModel, ValidationError, field_validator

router = APIRouter()

class UserPreferences(BaseModel):
    theme: str
    language: str
    notifications: bool

    @field_validator('theme')
    @classmethod
    def check_theme(cls, value: str) -> str:
        allowed_themes = ['light', 'dark']
        if value not in allowed_themes:
            raise ValueError('Invalid theme selection')
        return value

@router.post("/load-prefs")
async def load_prefs(data: str):
    try:
        raw_data = json.loads(data)
        preferences = UserPreferences(**raw_data)
        return preferences.model_dump()
    except (json.JSONDecodeError, ValidationError):
        raise HTTPException(status_code=400, detail="Datos invalidos")
