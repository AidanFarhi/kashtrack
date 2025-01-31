#!/bin/bash

DB_FILE_NAME="db/expense.db"
EC2_USER="ec2-user"
EC2_HOST="ec2-3-88-140-204.compute-1.amazonaws.com"
REMOTE_DIR="/home/ec2-user/app/db/"
SSH_KEY="./kashtrack-key-pair.pem"

scp -i $SSH_KEY $DB_FILE_NAME $EC2_USER@$EC2_HOST:$REMOTE_DIR