# Model training script

import mlflow
import numpy as np
from sklearn.linear_model import LogisticRegression

def train_model():
    # Sample data
    X = np.array([[1, 1], [1, 2], [2, 2], [2, 3]])
    y = np.array([0, 0, 1, 1])
    
    with mlflow.start_run():
        model = LogisticRegression()
        model.fit(X, y)
        accuracy = model.score(X, y)
        
        # Log metrics and model
        mlflow.log_metric("accuracy", accuracy)
        mlflow.sklearn.log_model(model, "model")
        print(f"Model trained with accuracy: {accuracy:.2f}")

if __name__ == "__main__":
    train_model()