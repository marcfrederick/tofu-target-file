resource "aws_s3_bucket" "first" {
  bucket = "first-bucket"
}

resource "aws_s3_bucket" "second" {
  bucket = "second-bucket"
}

resource "aws_iam_role" "example" {
  name               = "example-role"
  assume_role_policy = "{}"
}
