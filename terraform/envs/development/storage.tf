resource "aws_instance" "ec2-storage" {

  ami                         = var.ami_id
  instance_type               = var.instance_type
  subnet_id                   = aws_subnet.subnet-storage.id
  associate_public_ip_address = true
  key_name                    = var.ssh_public_key_name
  security_groups             = ["${aws_security_group.sg-storage.id}", "${aws_security_group.sg-ssh.id}"]

  tags = merge(
    var.tags,
    {
      Name = "ec2-storage-${var.project}-${var.environment}"
    }
  )

}

resource "aws_eip" "eip-ec2-storage" {
  instance = aws_instance.ec2-storage.id
  vpc      = true
}

resource "aws_security_group" "sg-storage" {
  vpc_id = aws_vpc.vpc.id
}

resource "aws_security_group_rule" "storage_inbound_for_api" {
  type                     = "ingress"
  from_port                = 4080
  to_port                  = 4080
  protocol                 = "tcp"
  source_security_group_id = aws_security_group.sg-api.id
  security_group_id        = aws_security_group.sg-storage.id
}

resource "aws_security_group_rule" "storage_inbound_for_indexer" {
  type                     = "ingress"
  from_port                = 4080
  to_port                  = 4080
  protocol                 = "tcp"
  source_security_group_id = aws_security_group.sg-indexer.id
  security_group_id        = aws_security_group.sg-storage.id
}
