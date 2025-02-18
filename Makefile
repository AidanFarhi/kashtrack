
.PHONY: open_firewall
open_firewall:
	@./scripts/open_firewall.sh

.PHONY: close_firewall
close_firewall:
	@./scripts/close_firewall.sh

.PHONY: startserver
startserver:
	@./scripts/open_firewall.sh
	@./scripts/start_server.sh
	@./scripts/close_firewall.sh

.PHONY: stopserver
stopserver:
	@./scripts/open_firewall.sh
	@./scripts/stop_server.sh
	@./scripts/close_firewall.sh

.PHONY: deploy
deploy:
	@./scripts/open_firewall.sh
	@./scripts/start_server.sh
	@./scripts/deploy.sh
	@./scripts/close_firewall.sh
