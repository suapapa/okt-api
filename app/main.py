from fastapi import FastAPI, Request
from pydantic import BaseModel
from starlette.middleware.base import BaseHTTPMiddleware
from starlette.responses import JSONResponse

import openkoreantext as okt
import os

# Check for required environment variable
okt_token = os.getenv("OKT_TOKEN")
if not okt_token:
    raise ValueError("OKT_TOKEN environment variable is required")

class TokenMiddleware(BaseHTTPMiddleware):
    async def dispatch(self, request: Request, call_next):
        token = request.headers.get("Authorization")
        if not token or token != f"Bearer {okt_token}":
            return JSONResponse(status_code=401, content={"detail": "Invalid or missing token"})
        return await call_next(request)

app = FastAPI(root_path=os.getenv("ROOT_PATH", "/okt"))
app.add_middleware(TokenMiddleware)

class TextRequest(BaseModel):
    text: str

@app.post("/normalize")
async def normalize_text(req: TextRequest):
    normalized = okt.normalize(req.text)
    return {"normalized": f"{normalized}"}

@app.post("/phrases")
async def extract_phrases(req: Request, text_req: TextRequest):
    include_hashtags = req.query_params.get('include_hashtags', 'false').lower() == 'true'
    phrases = okt.phrases(text_req.text, include_hashtags=include_hashtags)
    formatted_phrases = [f"{phrase}" for phrase in phrases]
    return {"phrases": formatted_phrases}
