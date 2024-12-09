# K8s webapp

kubectl apply -f k8s-webapp

kubectl port-forward $(kubectl get pods -l app=hello-world-label -o jsonpath="{.items[0].metadata.name}") 8000:80