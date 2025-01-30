#!/bin/bash

APP_NAME="kashtrack"
GO_FILE="app.go"
LOG_FILE="app.log"
BINARY_NAME="kashtrack"
WEB_DIR="web"
EC2_USER="ec2-user"
EC2_HOST="ec2-98-84-140-246.compute-1.amazonaws.com"
GOOS="linux"
GOARCH="amd64"
REMOTE_DIR="/home/ec2-user/app"
SSH_KEY="./kashtrack-key-pair.pem"
LOG_PATH=$REMOTE_DIR/$LOG_FILE 
CRON_CMD="*/1 * * * * tail -n 5 $LOG_PATH > $REMOTE_DIR/tmp_fl && mv $REMOTE_DIR/tmp_fl $LOG_PATH"

echo "Building Go binary..."
GOOS=$GOOS GOARCH=$GOARCH go build -o $BINARY_NAME $GO_FILE

echo "Killing the old process..."
ssh -i $SSH_KEY $EC2_USER@$EC2_HOST sudo pkill -f $BINARY_NAME

echo "Creating log file..."
ssh -i $SSH_KEY $EC2_USER@$EC2_HOST touch $REMOTE_DIR/$LOG_FILE

echo "Copying binary..."
scp -i $SSH_KEY $BINARY_NAME $EC2_USER@$EC2_HOST:$REMOTE_DIR/

echo "Copying the web files..."
scp -r -i $SSH_KEY $WEB_DIR $EC2_USER@$EC2_HOST:$REMOTE_DIR/

echo "Starting application..."
ssh -i "$SSH_KEY" "$EC2_USER@$EC2_HOST" <<EOF
    cd $REMOTE_DIR
    sudo nohup ./$BINARY_NAME $LOG_FILE > /dev/null 2>&1 &
    (crontab -l 2>/dev/null; echo "$CRON_CMD") | crontab -
EOF

echo "Done."
