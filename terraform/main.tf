provider "aws" {
  region = "us-east-1"
}

resource "aws_s3_bucket" "public_bucket" {
  bucket = "my-insecure-bucket-12345"
  acl    = "public-read"
}
