NAME := irlrobot/leadership
#TAG := $$(git rev-parse --short HEAD)
TAG := latest
IMG := ${NAME}:${TAG}
#LATEST := ${NAME}:latest
 
build:
	@docker build -t ${IMG} .
	#@docker tag ${IMG} ${LATEST}
 
push:
	@docker push ${NAME}
 
login:
	@docker log -u ${DOCKER_USER} -p ${DOCKER_PASS}

run:
	@docker run -it --rm -p 8080:8080 irlrobot/leadership