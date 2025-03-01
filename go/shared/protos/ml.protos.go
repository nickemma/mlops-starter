
// shared/protos/ml.proto
service MLService {
  rpc Predict(PredictRequest) returns (PredictResponse);
}

message PredictRequest {
  repeated float features = 1;
}

message PredictResponse {
  float prediction = 1;
  string model_version = 2;
}