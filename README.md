# Microservices

```mermaid
graph LR
    Internet -->|Request| messenger-gateway
    messenger-gateway -->|Route| messenger-vs
    messenger-vs -->|Forward| messenger-service
    messenger-service -->|Proxy| messenger-pod
```

Для Canary deployment эти ресурсы используются так:
1) Pods: разные версии приложения
2) Services: доступ к разным версиям
3) Gateway: входная точка в кластер
4) VirtualService: правила распределения трафика между версиями

