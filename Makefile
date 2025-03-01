# Train model (Python)
cd python/training && python train.py

# Build Go services
cd go && make build

# Run combined system
docker-compose up --build