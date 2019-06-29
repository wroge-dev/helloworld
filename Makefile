push-dockerhub:
	GOOS=linux go build .
	docker build -t wroge/helloworld:$(version) .
	docker push wroge/helloworld:$(version)

push-github:
	git add .
	git commit -m "update"
	git push

delete-kubernetes:
	kubectl delete deployment,service helloworld

delete-swarm:
	docker service rm helloworld_api

new-kubernetes:
	kubectl create -f deployment.json
	kubectl create -f service.json

new-swarm:
	docker stack deploy -c docker-compose.yml helloworld