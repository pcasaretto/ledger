provider "aws" {
  region = local.region
}

variable "rds_size" {
  default = "db.m5.xlarge"
}

variable "name" {}
variable "region" {}

variable "test_upgrade_to_latest_src" {
  default = "v1.4.0"
}

locals {
  name   = var.name
  region = var.region
}