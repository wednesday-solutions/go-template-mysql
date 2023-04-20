#!/bin/bash

source ./vitess/scripts/utils.sh

minikube start --kubernetes-version=v1.24.0 --cpus=4 --memory=8000 --disk-size=30g
killall kubectl
kubectl apply -f ./vitess/local/operator.yaml
checkPodStatusWithTimeout "vitess-operator(.*)1/1(.*)Running(.*)"

kubectl apply -f ./vitess/local/01-initial-cluster.yaml
checkPodStatusWithTimeout "vitess-cluster-useast1-vtctld(.*)1/1(.*)Running(.*)"
checkPodStatusWithTimeout "vitess-cluster-useast1-vtgate(.*)1/1(.*)Running(.*)"
checkPodStatusWithTimeout "vitess-cluster-etcd(.*)1/1(.*)Running(.*)" 3
checkPodStatusWithTimeout "vitess-cluster-vttablet-useast1(.*)3/3(.*)Running(.*)" 2

echo "Found all the required components, please wait..."

sleep 10

echo "Forwarding ports..."
sh ./vitess/scripts/pf.sh
