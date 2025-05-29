provider "aws" {
  region = "us-east-1"
}

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }

  required_version = ">= 1.3.0"
}

resource "aws_key_pair" "rafe_tf_key" {
  key_name   = "rafe-tf-key"
  public_key = file("~/.ssh/rafe-tf-key.pub")
  tags = {
    Name = "rafe-tf-key"
  }
}

resource "aws_security_group" "refe_terraform_sg" {
  name        = "refe-terraform-sg"
  description = "Allow SSH from my IP and HTTP access"
  vpc_id      = data.aws_vpc.default.id

  ingress {
    description = "SSH from my IP"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["82.169.184.68/32"]
  }

  ingress {
    description = "Allow HTTP"
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "refe-terraform-sg"
  }
}

data "aws_vpc" "default" {
  default = true
}

resource "aws_instance" "example" {
  ami                    = "ami-0953476d60561c955" // Amazon Linux 2023 AMI (HVM)
  instance_type          = "t2.micro"
  key_name               = aws_key_pair.rafe_tf_key.key_name
  vpc_security_group_ids = [aws_security_group.refe_terraform_sg.id]

user_data = <<-EOF
    #!/bin/bash
    set -e
    sudo dnf update -y
    sudo dnf install nginx -y
    sudo systemctl enable nginx
    sudo systemctl start nginx
EOF

  tags = {
    Name = "rafe-terraform-example"
  }
}

output "ec2_ssh_command" {
  value       = "ssh -i ~/.ssh/rafe-tf-key ec2-user@${aws_instance.example.public_ip}"
  description = "SSH command to connect to the EC2 instance"
}

# For testing
output "instance_id" {
  value = aws_instance.example.id
}

output "public_ip" {
    value       = aws_instance.example.public_ip
    description = "The public IP address of the EC2 instance"
}

output "instance_state" {
    value       = aws_instance.example.instance_state
    description = "The state of the EC2 instance"
}
