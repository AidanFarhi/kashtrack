
startserver:
	@./scripts/start_server.sh

stopserver:
	@./scripts/stop_server.sh

deploy: startserver
	@./scripts/deploy.sh
