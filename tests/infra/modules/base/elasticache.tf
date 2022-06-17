resource "aws_elasticache_subnet_group" "this" {
  name       = local.name
  subnet_ids = module.vpc.public_subnets
}

resource "aws_elasticache_parameter_group" "this" {
  name   = local.name
  family = "redis6.x"
}

resource "aws_elasticache_replication_group" "this" {
  availability_zones            = ["${local.region}a", "${local.region}b", "${local.region}c"]
  replication_group_id          = local.name
  replication_group_description = local.name
  node_type                     = "cache.t4g.micro"
  number_cache_clusters         = 3
  parameter_group_name          = aws_elasticache_parameter_group.this.name
  subnet_group_name             = aws_elasticache_subnet_group.this.name
  security_group_ids            = [module.security_group.security_group_id]
  port                          = 6379

  at_rest_encryption_enabled = true
  transit_encryption_enabled = true
  multi_az_enabled           = true
  automatic_failover_enabled = true
}