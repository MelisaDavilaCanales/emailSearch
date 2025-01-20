resource "aws_instance" "ec2-client" {

  ami                         = var.ami_id
  instance_type               = var.instance_type
  subnet_id                   = module.vpc.public_subnets[0]
  vpc_security_group_ids      = [module.sg-client.security_group_id]
  associate_public_ip_address = true
  key_name                    = var.ssh_public_key_name

  tags = merge(
    var.tags,
    {
      name = "ec2-client-${var.project}-${var.environment}"
    }
  )

}

module "sg-client" {
  source = "terraform-aws-modules/security-group/aws"

  name        = "sg_client_${var.project}_${var.environment}"
  description = "Security group for client"
  vpc_id      = module.vpc.vpc_id

  # Ingress rules
  ingress_cidr_blocks = ["0.0.0.0/0"]   # Allow all public IPs (for web access)
  ingress_rules       = ["http-80-tcp"] # Allow HTTP traffic on port 80

  # Egress rules (default allows all outbound traffic)
  egress_rules = ["all-all"]

  tags = var.tags
}