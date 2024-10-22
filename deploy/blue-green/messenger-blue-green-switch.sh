#!/bin/bash

function switch_to_green() {
    kubectl patch service messenger-svc -p '{"spec":{"selector":{"version":"green"}}}'
    echo "Switched traffic to green deployment"
}

function switch_to_blue() {
    kubectl patch service messenger-svc -p '{"spec":{"selector":{"version":"blue"}}}'
    echo "Switched traffic to blue deployment"
}

case "$1" in
    "switch_to_green")
        switch_to_green
        ;;
    "switch_to_blue")
        switch_to_blue
        ;;
    *)
        echo "Usage: $0 {switch_to_green|switch_to_blue}"
        exit 1
        ;;
esac