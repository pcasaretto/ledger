module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "~> 3.0"

  name = local.name
  cidr = "192.168.0.0/18"

  azs              = ["${local.region}a", "${local.region}b", "${local.region}c"]
  public_subnets   = ["192.168.0.0/24", "192.168.1.0/24", "192.168.2.0/24"]
  private_subnets  = ["192.168.3.0/24", "192.168.4.0/24", "192.168.5.0/24"]
  database_subnets = ["192.168.7.0/24", "192.168.8.0/24", "192.168.9.0/24"]

  create_database_subnet_group       = true
  create_database_subnet_route_table = true
}

module "security_group" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "~> 4.0"

  name        = local.name
  description = "PostgreSQL security group"
  vpc_id      = module.vpc.vpc_id

  # ingress
  ingress_with_cidr_blocks = [
    {
      from_port   = 5432
      to_port     = 5432
      protocol    = "tcp"
      description = "PostgreSQL access from within VPC"
      cidr_blocks = module.vpc.vpc_cidr_block
    },
    {
      from_port   = 6379
      to_port     = 6379
      protocol    = "tcp"
      description = "Redis access from within VPC"
      cidr_blocks = module.vpc.vpc_cidr_block
    }
  ]

  egress_with_cidr_blocks = [
    {
      from_port   = 0
      to_port     = 0
      protocol    = -1
      description = "Allow all outbound traffic"
      cidr_blocks = "0.0.0.0/0"
  }]
}
