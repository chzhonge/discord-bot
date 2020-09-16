build-docker-image:
	(cd src; docker build -t asia.gcr.io/zhong-discord-bot/discord-bot:latest -f dockerfile .)
push-docker-image:
	docker push asia.gcr.io/zhong-discord-bot/discord-bot:latest
deploy-gke-conf:
	kubectl apply -f gke-secert.yaml -f gke-deploy.yaml
deploy: build-docker-image push-docker-image deploy-gke-conf
remove-none-image:
	docker images | grep 'none' | awk {'print $3'} | xargs -n 1 docker rmi -f
post-deploy: remove-none-image
