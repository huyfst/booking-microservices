# Booking Microservices

This project leverage Golang, Grpc, RabbitMQ, Elasticsearch... to build a modern microservice architect 

## Create cluster local
```
kind create cluster --config index.yaml 
```

## Install 
```
Elasticsearch: 
  - helm repo add elastic https://helm.elastic.co
  - helm install elasticsearch elastic/elasticsearch
```

```
Jaeger:
  - helm repo add jaegertracing https://jaegertracing.github.io/helm-charts
  - helm install jaeger jaegertracing/jaeger
```