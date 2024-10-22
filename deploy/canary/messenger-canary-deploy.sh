kubectl apply -f ./deploy/canary/deployments.yaml
kubectl apply -f ./deploy/canary/services.yaml
kubectl apply -f ./deploy/canary/istio-gateway.yaml
kubectl apply -f ./deploy/canary/virtual-service.yaml
kubectl apply -f ./deploy/canary/destination-rule.yaml
kubectl apply -f ./deploy/canary/monitoring.yaml

./deploy/canary/canary-control.sh 10

# Testing the deployment
# Monitoring the deployment

kubectl logs -l app=messenger,version=canary
kubectl get pods -l app=messenger

sleep 30

./deploy/canary/canary-control.sh 30

sleep 30

./deploy/canary/canary-control.sh 50

sleep 30

./deploy/canary/canary-control.sh 100

curl -H "x-canary: true" http://$GATEWAY_URL/

# Rolling back the deployment
./deploy/canary/canary-control.sh 0

kubectl port-forward svc/prometheus-server 9090:9090

kubectl port-forward svc/grafana 3000:3000
