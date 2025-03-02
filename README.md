<!-- Project documentation -->

# MLOps Starter Kit

A minimal hybrid Go/Python MLOps implementation demonstrating:

- Python model training with MLflow
- Go API gateway
- Kubernetes deployment
- Basic CI/CD workflow

## Project Structure

## Getting Started

chmod +x scripts/setup_mlflow.sh
./scripts/setup_mlflow.sh

### Prerequisites

- Go 1.21+
- Python 3.9+
- Docker
- Kubernetes cluster (Minikube recommended)
- MLflow tracking server

### Quick Start

1. Train model:

```bash
make train
```

2. Build and run Go API:

```bash
make build-go && ./bin/mlops-api
```

3. Test API:
   Test End-to-End

```bash
curl -X POST http://localhost:8080/predict \
  -H "Content-Type: application/json" \
  -d '{"features": [2.0, 3.0]}'
```

4. Deploy to Kubernetes:

```bash
make docker-build deploy
```

### CI/CD Pipeline

GitHub Actions workflow included for:

```bash
# cd.yml
if: ${{ env.DEPLOY_ENABLED == 'true' }}
```

- Automated testing
- Docker image building
- Kubernetes deployment

## Multi-Environment Support

```bash
strategy:
  matrix:
    environment: [staging, production]
```

- **Development**: Local development environment
- **Staging**: Kubernetes cluster with staging namespace
- **Production**: Kubernetes cluster with production namespace

## Monitoring

Basic Prometheus/Grafana monitoring setup included in /deployments/monitoring

---

### **Verification Workflow**

1. **Train Model**

```bash
# Start MLflow server
mlflow server --host 0.0.0.0 --port 5000

# In another terminal
make train
```

2. **Test Go API**

```bash
# Build and run Go API
make build-go
./bin/mlops-api

# Test endpoint
curl -X POST http://localhost:8080/predict \
  -H "Content-Type: application/json" \
  -d '{"features": [2.0, 3.0]}'

# Output: {"prediction": 5}
```

3. **Deploy to Kubernetes**

```bash
# Build Docker images
make docker-build

# Deploy
kubectl apply -f deployments/

# Check status
kubectl get pods
```

4. **CI/CD Pipeline**

```bash
# Automated testing
# Docker image building
# Kubernetes deployment
```

5. **Monitoring**

```bash
# Prometheus/Grafana monitoring
```

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
