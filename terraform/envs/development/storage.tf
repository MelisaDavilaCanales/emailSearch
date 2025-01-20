resource "aws_instance" "ec2-storage" {

  ami                         = var.ami_id
  instance_type               = var.instance_type
  subnet_id                   = module.vpc.public_subnets[2]
  vpc_security_group_ids      = [module.sg-storage.security_group_id]
  associate_public_ip_address = true
  key_name                    = var.ssh_public_key_name

  tags = merge(
    var.tags,
    {
      name = "ec2-storage-${var.project}-${var.environment}"
    }
  )

}

module "sg-storage" {
  source = "terraform-aws-modules/security-group/aws"

  name        = "sg_storage_${var.project}_${var.environment}"
  description = "Security group for Storage"
  vpc_id      = module.vpc.vpc_id

  # Ingress rules for Storage
  ingress_with_cidr_blocks = [
    {
      from_port   = 4080
      to_port     = 4080
      protocol    = "tcp"
      description = "Allows connection to the Storage API"
      cidr_blocks = "10.0.0.0/16" # Allow access only from API and other internal services
    },
  ]

  # Egress rules (default allows all outbound traffic)
  egress_rules = ["all-all"]

  tags = var.tags
}
