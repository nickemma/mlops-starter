# Python Service
FROM python:3.10-slim
COPY python/requirements.txt .
RUN pip install -r requirements.txt
COPY python /app
CMD ["uvicorn", "serving.main:app"]

# Go Service
FROM golang:1.21-alpine
COPY go /src
RUN cd /src && go build -o /app
CMD ["/app"]