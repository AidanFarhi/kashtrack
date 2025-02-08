#!/bin/bash

ssh_key=./kashtrack-key-pair.pem

echo "setting secrets"
source ./scripts/init_secrets.sh

echo "copying db"
scp -i $ssh_key db/expense.db ec2-user@kash-track.com:/home/ec2-user/app/db/expense.db

echo "done"