# src/python/routes/users.py
# PASO 21: SQL Injection — consultas parametrizadas con sqlite3

# CODIGO SEGURO
import sqlite3
from fastapi import APIRouter

router = APIRouter()

@router.get("/users")
async def get_user(username: str):
    conn = sqlite3.connect(":memory:")
    # El ? es un placeholder: sqlite3 serializa el valor de forma segura
    # El input del usuario NUNCA toca el texto de la query SQL
    cursor = conn.execute(
        "SELECT id, username, email FROM users WHERE username = ?",
        (username,)
    )
    rows = cursor.fetchall()
    conn.close()
    return {"users": rows}
