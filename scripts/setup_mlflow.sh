# MLflow initialization
#!/bin/bash

# Initialize MLflow
pip install mlflow
mlflow server \
    --backend-store-uri sqlite:///mlflow.db \
    --default-artifact-root s3://mlflow-artifacts/ \
    --host 0.0.0.0 \
    --port 5000 &

# Register model
python -c "
import mlflow
from sklearn.linear_model import LogisticRegression

with mlflow.start_run() as run:
    model = LogisticRegression()
    X = [[1,1], [1,2], [2,2], [2,3]]
    y = [0,0,1,1]
    model.fit(X, y)
    mlflow.sklearn.log_model(model, 'model', registered_model_name='demo_model')
"