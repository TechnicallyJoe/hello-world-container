# HELLO-WORLD-CONTAINER

## build
docker build . --tag technicallyjoe/hello-world-container:x.y.z

## push
docker push technicallyjoe/hello-world-container:x.y.z

## run
docker run -e HWC_SLEEPTIMER=5 hello-world-container:latest