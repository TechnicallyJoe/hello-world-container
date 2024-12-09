VERSION=azfunc-0.1.1
IMGNAME=technicallyjoe/hello-world-container

buildandpush:
	docker build . --tag ${IMGNAME}:${VERSION}
	docker push ${IMGNAME}:${VERSION}
	docker image rm ${IMGNAME}:${VERSION}