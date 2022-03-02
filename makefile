VERSION=0.0.0
IMGNAME=technicallyjoe/hello-world-container

buildandpush:
	docker build . --tag ${IMGNAME}:${VERSION}
	docker push ${IMGNAME}:${VERSION}
	docker image rm ${IMGNAME}:${VERSION}