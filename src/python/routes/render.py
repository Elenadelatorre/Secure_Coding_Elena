# src/python/routes/render.py
# PASO 3: Server-Side Template Injection (SSTI) — template fijo con autoescaping habilitado

# CODIGO SEGURO
from jinja2 import Environment, select_autoescape
from fastapi import APIRouter

router = APIRouter()

# El template es una constante definida por el desarrollador, nunca por el usuario
GREETING_TEMPLATE = "Hola {{ name }}!"

env = Environment(autoescape=select_autoescape(["html", "xml"]))

@router.get("/greet")
async def greet(name: str):
    template = env.from_string(GREETING_TEMPLATE)
    return {"message": template.render(name=name)}  # name es un dato, no sintaxis
