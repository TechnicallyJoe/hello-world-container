# K8s webapp with sealed secrets helm

kubectl apply -f k8s-webapp-sealed-secrets-helm

kubectl port-forward svc/svc/hello-world-container-app-hello-world-service 8000:80

## Requirements

### bitnami/sealed-secrets
https://github.com/bitnami-labs/sealed-secrets

sealed secrets provide a way to encrypt values so that you'll be able to put them into git repositories. These can only be decrypted by the kubernetes cluster.

```
helm repo add sealed-secrets https://bitnami-labs.github.io/sealed-secrets
helm repo update
helm upgrade --install sealed-secrets bitnami-labs.github.io/sealed-secrets

```

### bakito/sealed-secrets-web
https://github.com/bakito/sealed-secrets-web

This provides a web interface for sealed secrets.

--format=yaml changes the input from json to yaml (Preference)
--disable-load-secrets disables the decrypt feature. It's meant to be one-way encryption so i opted to disable this.

```
helm repo add bakito https://bakito.github.io/helm-charts
helm repo update
helm upgrade --install sealed-secrets-web bakito/sealed-secrets-web --set "image.args={--format=yaml,--disable-load-secrets}"

kubectl port-forward svc/hello-world-container-app-sealed-secrets-web 8001:80
```