from fastapi import FastAPI, Request, APIRouter
from pydantic import BaseModel
from starlette.middleware.base import BaseHTTPMiddleware
from starlette.responses import JSONResponse

import openkoreantext as okt
import os

# Check for required environment variable
with open("/secret/token", "r") as token_file:
    okt_token = token_file.read().strip()
if not okt_token:
    raise ValueError("Token file is empty or missing")

class TokenMiddleware(BaseHTTPMiddleware):
    async def dispatch(self, request: Request, call_next):
        # Skip token validation for GET requests
        if request.method == "GET":
            return await call_next(request)
            
        token = request.headers.get("Authorization")
        if not token or token != f"Bearer {okt_token}":
            return JSONResponse(status_code=401, content={"detail": "Invalid or missing token"})
        return await call_next(request)

app = FastAPI(root_path=os.getenv("ROOT_PATH", "/okt"))
app.add_middleware(TokenMiddleware)

# Create API router for v1
v1_router = APIRouter(prefix="/v1")

class TextRequest(BaseModel):
    text: str

@v1_router.post("/normalize")
async def normalize_text(req: TextRequest):
    normalized = okt.normalize(req.text)
    return {"normalized": f"{normalized}"}

@v1_router.post("/phrases")
async def extract_phrases(req: Request, text_req: TextRequest):
    include_hashtags = req.query_params.get('include_hashtags', 'false').lower() == 'true'
    phrases = okt.phrases(text_req.text, include_hashtags=include_hashtags)
    formatted_phrases = [f"{phrase}" for phrase in phrases]
    return {"phrases": formatted_phrases}

# Include the v1 router
app.include_router(v1_router)
