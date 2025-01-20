resource "aws_instance" "ec2-indexer" {

  ami                         = var.ami_id
  instance_type               = var.instance_type
  subnet_id                   = module.vpc.public_subnets[3]
  vpc_security_group_ids      = [module.sg-indexer.security_group_id]
  associate_public_ip_address = true
  key_name                    = var.ssh_public_key_name

  tags = merge(
    var.tags,
    {
      name = "ec2-indexer-${var.project}-${var.environment}"
    }
  )
}


module "sg-indexer" {
  source = "terraform-aws-modules/security-group/aws"

  name        = "sg_indexer_${var.project}_${var.environment}"
  description = "Security group for Indexer"
  vpc_id      = module.vpc.vpc_id

  # Ingress rules: Typically not needed for the indexer.
  ingress_rules       = []
  ingress_cidr_blocks = []

  # Custom rules for Storage 
  egress_with_cidr_blocks = [
    {
      from_port   = 4080
      to_port     = 4080
      protocol    = "tcp"
      description = "Allow indexer to connect to Storage on port 4080"
      cidr_blocks = "10.0.0.0/16" # Allow only local VPC IPs
    },
  ]

}
