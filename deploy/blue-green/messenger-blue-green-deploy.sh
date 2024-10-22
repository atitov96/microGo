kubectl apply -f ./deploy/blue-green/messenger-green-deployment.yaml

kubectl rollout status deployment/messenger-blue
kubectl rollout status deployment/messenger-green

kubectl get service messenger-svc -o jsonpath='{.spec.selector.version}'

# Testing the deployment
# Monitoring the deployment

./deploy/blue-green/messenger-blue-green-switch.sh switch_to_green

# Rolling back the deployment
./deploy/blue-green/messenger-blue-green-switch.sh switch_to_blue

kubectl patch virtualservice messenger-vs --type=json \
  -p='[{"op": "replace", "path": "/spec/http/0/route/0/weight", "value": 50},
       {"op": "replace", "path": "/spec/http/0/route/1/weight", "value": 50}]'

kubectl patch virtualservice messenger-vs --type=json \
  -p='[{"op": "replace", "path": "/spec/http/0/route/0/weight", "value": 0},
       {"op": "replace", "path": "/spec/http/0/route/1/weight", "value": 100}]'
