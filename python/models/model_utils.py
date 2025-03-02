# ML model wrapper
import mlflow
import numpy as np
from typing import List

class ModelWrapper:
    def __init__(self, model_uri: str):
        self.model = mlflow.pyfunc.load_model(model_uri)
        self.version = mlflow.get_run(self.model.metadata.run_id).data.tags.get("version", "1.0.0")
        
    def predict(self, features: List[float]) -> int:
        if len(features) != 2:
            raise ValueError("Exactly 2 features required")
            
        return int(self.model.predict(np.array([features])))[0]

    def get_model_info(self) -> dict:
        return {
            "version": self.version,
            "input_shape": 2,
            "output_classes": 2
        }