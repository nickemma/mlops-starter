# FastAPI model server
from fastapi import FastAPI
from models.model_utils import ModelWrapper
import os

app = FastAPI()
model = ModelWrapper(os.getenv("MLFLOW_MODEL_URI", "models:/demo_model/Production"))

@app.post("/predict")
async def predict(features: list[float]):
    return {
        "prediction": model.predict(features),
        "model_version": model.version
    }

@app.get("/health")
async def health():
    return {"status": "healthy"}