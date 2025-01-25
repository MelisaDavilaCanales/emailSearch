resource "aws_instance" "ec2-client" {

  ami                         = var.ami_id
  instance_type               = var.instance_type
  subnet_id                   = aws_subnet.subnet-client.id
  associate_public_ip_address = true
  key_name                    = var.ssh_public_key_name
  security_groups             = ["${aws_security_group.sg-client.id}", "${aws_security_group.sg-ssh.id}"]

  tags = merge(
    var.tags,
    {
      Name = "ec2-client-${var.project}-${var.environment}"
    }
  )

}

resource "aws_eip" "eip-ec2-client" {
  instance = aws_instance.ec2-client.id
  vpc      = true
}

resource "aws_security_group" "sg-client" {
  vpc_id = aws_vpc.vpc.id
}

resource "aws_security_group_rule" "client_http_inbound" {
  type              = "ingress"
  from_port         = 80
  to_port           = 80
  protocol          = "tcp"
  cidr_blocks       = ["0.0.0.0/0"]
  security_group_id = aws_security_group.sg-client.id
}

resource "aws_security_group_rule" "client_to_api" {
  type                     = "egress"
  from_port                = 8080
  to_port                  = 8080
  protocol                 = "tcp"
  source_security_group_id = aws_security_group.sg-client.id
  security_group_id        = aws_security_group.sg-api.id
}
