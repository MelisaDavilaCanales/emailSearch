resource "aws_internet_gateway" "gateway" {
  vpc_id = aws_vpc.vpc.id
}

resource "aws_vpc" "vpc" {
  cidr_block           = "172.16.0.0/16"
  enable_dns_support   = true
  enable_dns_hostnames = true

  tags = {
    Name = "vpc-${var.project}-${var.environment}"
  }
}

resource "aws_route_table" "route-table" {
  vpc_id = aws_vpc.vpc.id
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.gateway.id
  }
  tags = {
    Name = "rt-${var.project}-${var.environment}"
  }
}

resource "aws_route_table_association" "api-association" {
  subnet_id      = aws_subnet.subnet-api.id
  route_table_id = aws_route_table.route-table.id
}

resource "aws_route_table_association" "client-association" {
  subnet_id      = aws_subnet.subnet-client.id
  route_table_id = aws_route_table.route-table.id
}

resource "aws_route_table_association" "storage-association" {
  subnet_id      = aws_subnet.subnet-storage.id
  route_table_id = aws_route_table.route-table.id
}

resource "aws_route_table_association" "indexer-association" {
  subnet_id      = aws_subnet.subnet-indexer.id
  route_table_id = aws_route_table.route-table.id
}

resource "aws_subnet" "subnet-client" {
  vpc_id            = aws_vpc.vpc.id
  cidr_block        = "172.16.1.0/24"
  availability_zone = "${var.region}a"

  tags = {
    Name = "vpc-client-${var.project}-${var.environment}"
  }
}

resource "aws_subnet" "subnet-api" {
  vpc_id            = aws_vpc.vpc.id
  cidr_block        = "172.16.2.0/24"
  availability_zone = "${var.region}b"

  tags = {
    Name = "vpc-api-${var.project}-${var.environment}"
  }

}

resource "aws_subnet" "subnet-storage" {
  vpc_id            = aws_vpc.vpc.id
  cidr_block        = "172.16.3.0/24"
  availability_zone = "${var.region}c"

  tags = {
    Name = "vpc-storage-${var.project}-${var.environment}"
  }
}

resource "aws_subnet" "subnet-indexer" {
  vpc_id                  = aws_vpc.vpc.id
  cidr_block              = "172.16.4.0/24"
  map_public_ip_on_launch = true
  availability_zone       = "${var.region}d"

  tags = {
    Name = "vpc-indexer-${var.project}-${var.environment}"
  }
}