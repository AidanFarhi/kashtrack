app_name := kashtrack
go_file := app.go
log_file := app.log
binary_name := kashtrack
web_dir := web
user := ec2-user
host := ec2-50-17-69-239.compute-1.amazonaws.com
goos := linux
goarch := amd64
remote_dir := /home/ec2-user/app
ssh_key := ./kashtrack-key-pair.pem
log_path := $(remote_dir)/$(log_file )
cron_cmd := */1 * * * * tail -n 5 $(log_path) > $(remote_dir)/tmp_fl && mv $(remote_dir)/tmp_fl $(log_path)

deploy:

	echo "Building Go binary..."
	GOOS=$(goos) GOARCH=$(goarch)go build -o $(binary_name) $(go_file)

	echo "Killing the old process..."
	ssh -i $(ssh_key) $(user)@$(host)sudo pkill -f $(binary_name)

	echo "Creating log file..."
	ssh -i $(ssh_key) $(user)@$(host) touch $(remote_dir)/$(log_file)

	echo "Copying .env file..."
	scp -i $(ssh_key) .env_prod $(user)@$(host):$(remote_dir)/.env

	echo "Copying binary..."
	scp -i $(ssh_key) $(binary_name) $(user)@$(host):$(remote_dir)/

	echo "Copying the web files..."
	scp -r -i $(ssh_key) $(web_dir) $(user)@$(host):$(remote_dir)/

	echo "Starting application..."
	ssh -i "$(ssh_key)" "$(user)@$(host)" <<EOF
		cd $(remote_dir)
		sudo nohup ./$(binary_name) $(log_file) > /dev/null 2>&1 &
		(crontab -l 2>/dev/null; echo "$(cron_cmd)") | crontab -
	EOF

	echo "Done."
