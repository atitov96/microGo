#!/bin/bash

function set_canary_weight() {
    CANARY_WEIGHT=$1
    STABLE_WEIGHT=$((100-$CANARY_WEIGHT))

    cat <<EOF | kubectl apply -f -
apiVersion: networking.istio.io/v1beta1
kind: VirtualService
metadata:
  name: messenger
  namespace: default
spec:
  hosts:
  - "*"
  gateways:
  - messenger-gateway
  http:
  - match:
    - headers:
        x-canary:
          exact: "true"
    route:
    - destination:
        host: messenger-canary
        port:
          number: 80
  - route:
    - destination:
        host: messenger-stable
        port:
          number: 80
      weight: $STABLE_WEIGHT
    - destination:
        host: messenger-canary
        port:
          number: 80
      weight: $CANARY_WEIGHT
EOF

    echo "Set canary weight to $CANARY_WEIGHT%"
}

if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <weight-percentage>"
    echo "Example: $0 20 (routes 20% traffic to canary)"
    exit 1
fi

WEIGHT=$1
if ! [[ "$WEIGHT" =~ ^[0-9]+$ ]] || [ "$WEIGHT" -gt 100 ]; then
    echo "Error: Weight must be a number between 0 and 100"
    exit 1
fi

set_canary_weight $WEIGHT
