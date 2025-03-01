# Python: train.py
import mlflow
from sklearn.ensemble import RandomForestClassifier

def train_model(data):
    with mlflow.start_run():
        model = RandomForestClassifier()
        model.fit(data.X, data.y)
        mlflow.log_metric("accuracy", model.score(data.X_test, data.y_test))
        mlflow.sklearn.log_model(model, "model")