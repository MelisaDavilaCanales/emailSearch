module "vpc" {
  source = "terraform-aws-modules/vpc/aws"

  name = "vpc-${var.project}-${var.environment}"
  cidr = "10.0.0.0/16"

  # client, api, storage, indexer
  azs            = ["us-east-1a", "us-east-1b", "us-east-1c", "us-east-1d"]
  public_subnets = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24", "10.0.4.0/24"]

  enable_nat_gateway = false
  enable_vpn_gateway = false

  tags = var.tags
}
