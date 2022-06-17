resource "random_string" "random" {
  length  = 6
  special = false
  numeric = false
  upper   = false
}

module "base" {
  source = "./modules/base"

  name = local.name
  region = local.region
  test_upgrade_to_latest_src = "v1.4.0"
}

locals {
  name          = "ledger-bench-${random_string.random.result}"
  region        = "eu-west-1"
  rds_size      = "db.m5.xlarge"
  database_url  = "postgresql://${module.base.database_username}:${module.base.database_password}@${module.base.database_host}:5432/${module.base.database_name}"
  redis_url     = "redis://${module.base.redis_host}"
  params_ledger = "${local.params_ledger_sql} ${local.params_ledger_redis}"
  params_ledger_sql = "--storage.driver postgres --storage.postgres.conn_string ${local.database_url}"
  params_ledger_redis = "--lock-strategy redis --lock-strategy-redis-tls-enabled true --lock-strategy-redis-url ${local.redis_url}"
}

module "test_upgrade_database" {
  source = "./modules/instances"

  name          = "ledger-bench-${random_string.random.result}"
  region        = "eu-west-1"

  vpc_public_subnets = module.base.public_subnets
  vpc_vpc_id         = module.base.vpc_id
  security_group     = module.base.security_group
  instance_role      = module.base.instance_role
  params_ledger = local.params_ledger
  params_ledger_sql = local.params_ledger_sql
  params_ledger_redis = local.params_ledger_redis

  user_data = <<-EOT
  #!/bin/bash
  sudo yum install -y https://s3.amazonaws.com/ec2-downloads-windows/SSMAgent/latest/linux_amd64/amazon-ssm-agent.rpm
  sudo systemctl enable amazon-ssm-agent
  sudo systemctl start amazon-ssm-agent
  sudo rpm -i https://github.com/numary/ledger/releases/download/v${var.ledger_version}/numary_${var.ledger_version}_linux_amd64.rpm
  EOT
}

output "instance_id" {
  value = module.test_upgrade_database.instance_id
}

output "name" {
  value = local.name
}