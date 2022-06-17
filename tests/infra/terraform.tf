terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }
}

provider "aws" {
  region = local.region
}

variable "ledger_version" {
  #  default = "1.4.0"
  default = "1.5.0-rc.2"
}
