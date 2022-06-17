data "aws_ami" "amazon_linux" {
  most_recent = true
  owners      = ["amazon"]

  filter {
    name = "name"

    values = [
      "amzn2-ami-*-hvm-*-x86_64-gp2",
    ]
  }
}

resource "aws_ssm_document" "storage-scan" {
  name            = "${local.name}-storage-scan"
  document_format = "YAML"
  document_type   = "Command"

  content = <<DOC
schemaVersion: '1.2'
description: Launch Storage Scan for Numary
parameters: {}
runtimeConfig:
  'aws:runShellScript':
    properties:
      - id: '0.aws:runShellScript'
        runCommand:
          - numary storage scan ${var.params_ledger_sql}
DOC
}

resource "aws_ssm_document" "storage-list" {
  name            = "${local.name}-storage-list"
  document_format = "YAML"
  document_type   = "Command"

  content = <<DOC
schemaVersion: '1.2'
description: Launch Storage List for Numary
parameters: {}
runtimeConfig:
  'aws:runShellScript':
    properties:
      - id: '0.aws:runShellScript'
        runCommand:
          - numary storage list ${var.params_ledger_sql}
DOC
}

resource "aws_ssm_document" "storage-upgrade" {
  name            = "${local.name}-storage-upgrade"
  document_format = "YAML"
  document_type   = "Command"

  content = <<DOC
schemaVersion: '1.2'
description: Launch Storage Upgrade for Numary
parameters: {}
runtimeConfig:
  'aws:runShellScript':
    properties:
      - id: '0.aws:runShellScript'
        runCommand:
          - numary storage upgrade c6561474-ad8e-43f4-aea7-e16b7fb80df7 ${var.params_ledger_sql}
DOC
}

resource "aws_instance" "ledger" {
  ami                  = data.aws_ami.amazon_linux.id
  instance_type        = "c5.large"
  user_data_base64     = base64encode(var.user_data)
  subnet_id            = var.vpc_public_subnets[0]
  iam_instance_profile = var.instance_role
  vpc_security_group_ids      = var.security_group
  key_name             = "maxence"
}