module "rds" {
  source = "terraform-aws-modules/rds/aws"

  identifier = local.name

  create_db_option_group    = false
  create_db_parameter_group = false

  engine               = "postgres"
  engine_version       = "14.1"
  family               = "postgres14"
  major_engine_version = "14"
  instance_class       = var.rds_size

  snapshot_identifier = "arn:aws:rds:eu-west-1:955332203423:snapshot:ledger-${replace(var.test_upgrade_to_latest_src, ".", "-")}"

  allocated_storage = 200

  db_name  = "ledger"
  username = "ledger"
  port     = 5432

  db_subnet_group_name   = module.vpc.database_subnet_group
  vpc_security_group_ids = [module.security_group.security_group_id]

  maintenance_window      = "Mon:00:00-Mon:03:00"
  backup_window           = "03:00-06:00"
  backup_retention_period = 0
}