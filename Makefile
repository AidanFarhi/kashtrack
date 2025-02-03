app_name := kashtrack
go_file := app.go
log_file := app.log
binary_name := kashtrack
web_dir := web
user := ec2-user
server_name := kashtrack-server
goos := linux
goarch := amd64
remote_dir := /home/ec2-user/app
ssh_key := ./kashtrack-key-pair.pem
log_path := $(remote_dir)/$(log_file)
cron_cmd := */1 * * * * tail -n 5 $(log_path) > $(remote_dir)/tmp_fl && mv $(remote_dir)/tmp_fl $(log_path)

startserver:
	@. ./init_secrets.sh && \
	instance_id=$$( \
		aws ec2 describe-instances \
			--filters "Name=tag:Name,Values=$(server_name)" \
			--query "Reservations[*].Instances[*].InstanceId" \
			--output text \
	) && \
	aws ec2 start-instances --instance-ids $$instance_id && \
	aws ec2 wait instance-running --instance-ids $$instance_id && \
	echo "server running"

stopserver:
	@. ./init_secrets.sh && \
	instance_id=$$( \
		aws ec2 describe-instances \
			--filters "Name=tag:Name,Values=$(server_name)" \
			--query "Reservations[*].Instances[*].InstanceId" \
			--output text \
	) && \
	aws ec2 stop-instances --instance-ids $$instance_id && \
	aws ec2 wait instance-stopped --instance-ids $$instance_id && \
	echo "server stopped"

deploy:
	@. ./init_secrets.sh; \
	host_name=$$( \
		aws ec2 describe-instances \
		--filters "Name=tag:Name,Values=$(server_name)" \
		--query "Reservations[*].Instances[*].PublicDnsName" \
		--output text \
	); \
	GOOS=$(goos) GOARCH=$(goarch) go build -o $(binary_name) $(go_file); \
	ssh -i $(ssh_key) $(user)@$$host_name sudo pkill -f $(binary_name); \
	ssh -i $(ssh_key) $(user)@$$host_name touch $(remote_dir)/$(log_file); \
	scp -i $(ssh_key) .env_prod $(user)@$$host_name:$(remote_dir)/.env; \
	scp -i $(ssh_key) $(binary_name) $(user)@$$host_name:$(remote_dir)/; \
	scp -r -i $(ssh_key) $(web_dir) $(user)@$$host_name:$(remote_dir)/; \
	ssh -i $(ssh_key) $(user)@$$host_name "cd $(remote_dir) && sudo nohup ./$(binary_name) > /dev/null 2>&1 & disown"; \
	ssh -i $(ssh_key) $(user)@$$host_name "(crontab -l 2>/dev/null; echo \"$(cron_cmd)\") | crontab -"; \
	echo "deployment complete: $$host_name"
