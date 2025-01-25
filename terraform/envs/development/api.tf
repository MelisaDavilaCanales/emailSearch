resource "aws_instance" "ec2-api" {
  ami                         = var.ami_id
  instance_type               = var.instance_type
  subnet_id                   = aws_subnet.subnet-client.id
  associate_public_ip_address = true
  key_name                    = var.ssh_public_key_name
  security_groups             = ["${aws_security_group.sg-api.id}", "${aws_security_group.sg-ssh.id}"]

  tags = merge(
    var.tags,
    {
      Name = "ec2-api-${var.project}-${var.environment}"
    }
  )
}

resource "aws_eip" "eip-ec2-api" {
  instance = aws_instance.ec2-api.id
  vpc      = true
}

resource "aws_security_group" "sg-api" {
  vpc_id = aws_vpc.vpc.id
}

resource "aws_security_group_rule" "api_inbound" {
  type                     = "ingress"
  from_port                = 8080
  to_port                  = 8080
  protocol                 = "tcp"
  source_security_group_id = aws_security_group.sg-client.id
  security_group_id        = aws_security_group.sg-api.id
}

resource "aws_security_group_rule" "api_to_storage" {
  type                     = "egress"
  from_port                = 4080
  to_port                  = 4080
  protocol                 = "tcp"
  source_security_group_id = aws_security_group.sg-api.id
  security_group_id        = aws_security_group.sg-storage.id
}

