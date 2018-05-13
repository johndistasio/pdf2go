provider "aws" {
  version = "~> 1.18"
}

variable "deployment_bundle" {
  default = "deployment.zip"
}

data "aws_iam_policy_document" "lambda" {
  statement {
    actions = ["sts:AssumeRole"]

    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "lambda" {
  name               = "Lambda"
  assume_role_policy = "${data.aws_iam_policy_document.lambda.json}"
}

resource "aws_iam_role_policy_attachment" "lambda" {
  role       = "${aws_iam_role.lambda.name}"
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

resource "aws_lambda_function" "hello" {
  filename         = "${var.deployment_bundle}"
  source_code_hash = "${base64sha256(file(var.deployment_bundle))}"

  runtime = "go1.x"

  function_name = "Hello"
  handler       = "main"

  role = "${aws_iam_role.lambda.arn}"

  environment {
    variables = {
      NAME = "Terraform"
    }
  }
}
