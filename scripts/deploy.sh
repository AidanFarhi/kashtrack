#!/bin/bash

# s3_bucket=afarhidev-kashtrack
ssh_key=./kashtrack-key-pair.pem
# log_cleanup_cron_cmd=*/1 * * * * tail -n 5 $(log_path) > /home/ec2-user/app/tmp_fl && mv /home/ec2-user/app/tmp_fl $(log_path)
# db_backup_cron_cmd := 0 23 * * * aws s3 cp /home/ec2-user/app/db/$(db_file).db s3://$(s3_bucket)/db/$(db_file)_$(date +\%Y_\%m_\%d).db

echo "starting deployment"

echo "building binary"
GOOS=linux GOARCH=amd64 go build -o kashtrack app.go

echo "stopping app"
ssh -o StrictHostKeyChecking=no -i $ssh_key ec2-user@kash-track.com sudo -n pkill -f kashtrack || true

echo "creating log file"
ssh -i $ssh_key ec2-user@kash-track.com touch /home/ec2-user/app/app.log

echo "copying binary and static files"
scp -i $ssh_key .env_prod ec2-user@kash-track.com:/home/ec2-user/app/.env
scp -i $ssh_key kashtrack ec2-user@kash-track.com:/home/ec2-user/app/
scp -r -i $ssh_key web ec2-user@kash-track.com:/home/ec2-user/app/

echo "starting app"
ssh -i $ssh_key ec2-user@kash-track.com "cd /home/ec2-user/app && echo 'sudo -n ./kashtrack > /dev/null 2>&1 &' | at now 2>/dev/null"

# echo "removing old cron jobs"
# ssh -i $ssh_key ec2-user@kash-track.com crontab -r || true

# echo "starting cron jobs"
# ssh -i $ssh_key ec2-user@kash-track.com 'echo "$(log_cleanup_cron_cmd)" | crontab -'

rm kashtrack
echo "deployment complete"
