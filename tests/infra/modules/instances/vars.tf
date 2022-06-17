provider "aws" {
  region = local.region
}

variable "name" {}
variable "region" {}
variable "user_data" {}
variable "vpc_public_subnets" {}
variable "vpc_vpc_id" {}
variable "instance_role" {}
variable "security_group" {}
variable "params_ledger" {}
variable "params_ledger_sql" {}
variable "params_ledger_redis" {}

locals {
  name   = var.name
  region = var.region
}