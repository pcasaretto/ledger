#!/usr/bin/env bash

TF_NAME=$(terraform output name | sed -e 's/^"//' -e 's/"$//')
TF_INSTANCE=$(terraform output instance_id | sed -e 's/^"//' -e 's/"$//')

function launch_and_get {
  echo "Send command: numary ${1} ${2}"
  CMD_SCAN=$(aws ssm send-command --document-name "${TF_NAME}-${1}-${2}" --document-version "1" --targets Key=InstanceIds,Values=${TF_INSTANCE} --parameters '{}' --timeout-seconds 600 --max-concurrency "50" --max-errors "0" --region eu-west-1 | jq -r '.Command.CommandId')
  sleep 10
  CMD_SCAN_RESULT=$(aws ssm get-command-invocation --command-id "${CMD_SCAN}" --instance-id "${TF_INSTANCE}" | jq -r '.Status' )
  echo $CMD_SCAN_RESULT

  if [ "$CMD_SCAN_RESULT" != "Success" ]
  then
      echo "$CMD_SCAN_RESULT"
      terraform destroy -auto-approve
      exit 1
  fi
}

terraform init
terraform apply -auto-approve
sleep 60
launch_and_get storage scan
launch_and_get storage list
launch_and_get storage upgrade
