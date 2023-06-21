
Zakładam, że masz już skonfigurowane credentiale AWS.

Plik `provider.tf`:

```hcl
provider "aws" {
  region = "eu-west-1"
}
```

Plik `vpc.tf`:

```hcl
resource "aws_vpc" "example" {
  cidr_block = "10.0.0.0/16"
  tags = {
    Name = "example-vpc"
  }
}

resource "aws_subnet" "example" {
  vpc_id     = aws_vpc.example.id
  cidr_block = "10.0.1.0/24"
  tags = {
    Name = "example-subnet"
  }
}
```

Plik `security_group.tf`:

```hcl
resource "aws_security_group" "example" {
  name        = "example"
  description = "Example security group"
  vpc_id      = aws_vpc.example.id

  ingress {
    from_port   = 22
    to_port     = 22
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
    Name = "example"
  }
}
```

Plik `ec2.tf`:

```hcl
resource "aws_instance" "example" {
  ami           = "ami-0c94855ba95c574c8" // Amazon Linux 2 AMI
  instance_type = "t2.micro"

  vpc_security_group_ids = [aws_security_group.example.id]
  subnet_id              = aws_subnet.example.id

  tags = {
    Name = "example-instance"
  }
}
```

Plik `iam.tf`:

```hcl
resource "aws_iam_role" "example" {
  name = "example_role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "ec2.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "example" {
  name = "example_policy"
  role = aws_iam_role.example.id

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "ec2:Describe*"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}
EOF
}

resource "aws_iam_instance_profile" "example" {
  name = "example_profile"
  role = aws_iam_role.example.name
}

resource "aws_instance" "example_with_iam" {
  ami           = "ami-0c94855ba95c574c8" // Amazon Linux 2 AMI
  instance_type = "t2.micro"

  vpc_security_group_ids = [aws_security_group.example.id]
  subnet_id              = aws_subnet.example.id
  iam_instance

_profile   = aws_iam_instance_profile.example.name

  tags = {
    Name = "example-instance-with-iam"
  }
}
```

Zwróć uwagę, że powyższy kod zakłada, że jesteś w regionie `eu-west-1` i używasz Amazon Linux 2 AMI o ID `ami-0c94855ba95c574c8`. Będziesz musiał dostosować te wartości do swojego regionu i wybranej AMI.

Ważne: Powyższy kod jest tylko przykładem i może wymagać dostosowania do Twoich konkretnych wymagań i środowiska.
