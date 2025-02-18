#!/bin/bash

echo "setting secrets"
source ./scripts/init_secrets.sh

echo "getting instance ID"
instance_id=$( \
    aws ec2 describe-instances \
        --filters "Name=tag:Name,Values=kashtrack-server" \
        --query "Reservations[*].Instances[*].InstanceId" \
        --output text \
)

echo "retrieving security group ID"
security_group_id=$( \
    aws ec2 describe-instances \
        --instance-ids $instance_id \
        --query "Reservations[*].Instances[*].SecurityGroups[*].GroupId" \
        --output text \
)

echo "fetching current public IP"
my_ip=$(curl -s https://checkip.amazonaws.com)

echo "updating security group to allow SSH access from $my_ip"
aws ec2 authorize-security-group-ingress \
    --group-id $security_group_id \
    --protocol tcp \
    --port 22 \
    --cidr ${my_ip}/32 \
    > /dev/null 2>&1

echo "security group updated successfully"
