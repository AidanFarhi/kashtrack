#!/bin/bash

echo "setting secrets"
source ./scripts/init_secrets.sh

echo "getting instance id"
instance_id=$( \
    aws ec2 describe-instances \
        --filters "Name=tag:Name,Values=kashtrack-server" \
        --query "Reservations[*].Instances[*].InstanceId" \
        --output text \
)

echo "stopping instance"
aws ec2 stop-instances --instance-ids $instance_id > /dev/null 2>&1
aws ec2 wait instance-stopped --instance-ids $instance_id

echo "server stopped"