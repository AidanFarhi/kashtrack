
startserver:
	@./scripts/start_server.sh

stopserver:
	@./scripts/stop_server.sh

migratedb: startserver
	@./scripts/migrate_db.sh

deploy: startserver
	@./scripts/deploy.sh