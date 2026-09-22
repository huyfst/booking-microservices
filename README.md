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

```
Cert Manager: 
  - kubectl apply -f https://github.com/cert-manager/cert-manager/releases/latest/download/cert-manager.yaml
```

```
RabbitMQ Cluster Operator
  - kubectl apply -f https://github.com/rabbitmq/cluster-operator/releases/latest/download/cluster-operator.yml
```

```
Nginx Ingress Controller
  - kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/cloud/deploy.yaml
```

```
Redis Sentinel
  - helm repo add bitnami https://charts.bitnami.com/bitnami
  - helm install redis bitnami/redis --set auth.password=**** -f redis.yaml
```