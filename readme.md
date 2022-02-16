# HELLO-WORLD-CONTAINER

## docker

### build
docker build . --tag technicallyjoe/hello-world-container:x.y.z

### push
docker push technicallyjoe/hello-world-container:x.y.z

### run
docker run -e HWC_SLEEPTIMER=5 hello-world-container:latest

## helm

### build
helm package helm --debug

### install
helm install hello-world-container chart-0.1.0.tgz

### uninstall
helm uninstall hello-world-container chart-0.1.0.tgz
