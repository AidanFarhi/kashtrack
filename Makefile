
.PHONY: startserver
startserver:
	@./scripts/start_server.sh

.PHONY: stopserver
stopserver:
	@./scripts/stop_server.sh

.PHONY: migratedb
migratedb: startserver
	@./scripts/migrate_db.sh

.PHONY: deploy
deploy: startserver
	@./scripts/deploy.sh
