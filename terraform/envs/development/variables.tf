variable "access_key" {
  description = "AWS Access Key"
  type        = string
  sensitive   = true
}

variable "secret_key" {
  description = "AWS Secret Key"
  type        = string
  sensitive   = true
}

variable "region" {
  description = "AWS Region"
  default     = "us-east-1"
  type        = string
}

variable "project" {
  description = "the name project"
  default     = "msc" #MailScout
}

variable "environment" {
  description = "the enviroment to release"
  default     = "dev"
}

variable "ami_id" {
  description = "the ami id to use"
  default     = "ami-011899242bb902164" # Ubuntu 20.04 LTS // us-east-1
}

variable "instance_type" {
  description = "the instance type to use"
  default     = "t2.micro"
}

variable "ssh_public_key_name" {
  description = "Name of the SSH public key (type ED25519)"
}

variable "tags" {
  description = "all tags used"
  default = {
    environment = "dev"
    project     = "msc"
    owner       = "melisadavila2001@gmail.com"
  }
}

