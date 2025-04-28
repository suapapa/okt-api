from fastapi import FastAPI
from pydantic import BaseModel

import openkoreantext as okt

app = FastAPI()

class TextRequest(BaseModel):
    text: str

@app.post("/normalize")
async def normalize_text(req: TextRequest):
    normalized = okt.normalize(req.text)
    return {"normalized": normalized}

@app.post("/phrases")
async def extract_phrases(req: TextRequest):
    phrases = okt.phrases(req.text, include_hashtags=req.query_params.get('include_hashtags', 'false').lower() == 'true')
    return {"phrases": phrases}
