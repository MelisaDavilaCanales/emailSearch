resource "aws_instance" "ec2-indexer" {

  ami                         = var.ami_id
  instance_type               = var.instance_type
  subnet_id                   = aws_subnet.subnet-indexer.id
  associate_public_ip_address = true
  key_name                    = var.ssh_public_key_name
  security_groups             = ["${aws_security_group.sg-indexer.id}", "${aws_security_group.sg-ssh.id}"]

  tags = merge(
    var.tags,
    {
      Name = "ec2-indexer-${var.project}-${var.environment}"
    }
  )
}

resource "aws_eip" "eip-ec2-indexer" {
  instance = aws_instance.ec2-indexer.id
  vpc      = true
}

resource "aws_security_group" "sg-indexer" {
  vpc_id = aws_vpc.vpc.id
}

resource "aws_security_group_rule" "indexer_to_storage" {
  type                     = "egress"
  from_port                = 4080
  to_port                  = 4080
  protocol                 = "tcp"
  source_security_group_id = aws_security_group.sg-indexer.id
  security_group_id        = aws_security_group.sg-storage.id
}