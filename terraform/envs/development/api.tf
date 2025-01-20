resource "aws_instance" "ec2-api" {

  ami                         = var.ami_id
  instance_type               = var.instance_type
  subnet_id                   = module.vpc.public_subnets[1]
  vpc_security_group_ids      = [module.sg-api.security_group_id]
  associate_public_ip_address = true
  key_name                    = var.ssh_public_key_name

  tags = merge(
    var.tags,
    {
      name = "ec2-api-${var.project}-${var.environment}"
    }
  )

}

module "sg-api" {
  source = "terraform-aws-modules/security-group/aws"

  name        = "sg_api_${var.project}_${var.environment}"
  description = "Security group for API"
  vpc_id      = module.vpc.vpc_id

  # Ingress rules
  ingress_cidr_blocks = ["0.0.0.0/0"] # Allow HTTP traffic from the client 
  ingress_rules       = ["http-8080-tcp"]

  # Custom rules for Storage 
  egress_with_cidr_blocks = [
    {
      from_port   = 4080
      to_port     = 4080
      protocol    = "tcp"
      description = "Allow indexer to connect to Storage on port 4080"
      cidr_blocks = "10.0.0.0/16"
    },
  ]

  # Egress rules (default allows all outbound traffic)
  egress_rules = ["all-all"]

  tags = var.tags
}
