output "vpc_id" {
  value = module.vpc.vpc_id
}

output "public_subnets" {
  value = module.vpc.public_subnets
}

output "database_host" {
  value = module.rds.db_instance_address
}

output "database_password" {
  value = module.rds.db_instance_password
}

output "database_username" {
  value = module.rds.db_instance_username
}

output "database_name" {
  value = module.rds.db_instance_name
}

output "redis_host" {
  value = aws_elasticache_replication_group.this.primary_endpoint_address
}

output "instance_role" {
  value = aws_iam_instance_profile.test_profile.id
}

output "security_group" {
  value = [module.security_group.security_group_id]
}