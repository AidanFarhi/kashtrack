#!/bin/bash

APP_NAME="kashtrack"
GO_FILE="app.go"
BINARY_NAME="kashtrack"
WEB_DIR="web"
EC2_USER="ec2-user"
EC2_HOST="ec2-3-88-140-204.compute-1.amazonaws.com"
GOOS="linux"
GOARCH="amd64"
REMOTE_DIR="/home/ec2-user/app"
SSH_KEY="./kashtrack-key-pair.pem"

# Step 1: Build the Go binary
echo "Building Go binary..."
GOOS=$GOOS GOARCH=$GOARCH go build -o $BINARY_NAME $GO_FILE
if [ $? -ne 0 ]; then
    echo "Failed to build the Go binary."
    exit 1
fi

echo "Killing the old process..."
ssh -i $SSH_KEY $EC2_USER@$EC2_HOST << EOF
    pkill -f $BINARY_NAME
EOF

# Step 2: Copy the binary to the EC2 instance
echo "Copying binary to EC2 instance..."
scp -i $SSH_KEY $BINARY_NAME $EC2_USER@$EC2_HOST:$REMOTE_DIR/
echo "Copying the web files to the EC2 instance..."
scp -r -i $SSH_KEY $WEB_DIR $EC2_USER@$EC2_HOST:$REMOTE_DIR/
if [ $? -ne 0 ]; then
    echo "Failed to copy the binary and web files to the EC2 instance."
    exit 1
fi

# Step 4: SSH into the EC2 instance and restart the application
echo "Starting application on EC2 instance..."
ssh -i $SSH_KEY $EC2_USER@$EC2_HOST << EOF
    export KASHTRACK_APP_LOGS="app.log"
    cd $REMOTE_DIR
    sudo nohup ./$BINARY_NAME > \$KASHTRACK_APP_LOGS 2>&1 &
    echo "Application started successfully."
EOF

if [ $? -ne 0 ]; then
    echo "Failed to restart the application on the EC2 instance."
    exit 1
fi

echo "Deployment completed successfully."
