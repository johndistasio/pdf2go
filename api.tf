provider "aws" {
  version = "~> 1.18"
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

variable "deployment_bundle" {
  default = "deployment.zip"
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
      NAME = "Hello"
    }
  }
}

resource "aws_api_gateway_rest_api" "hello" {
  name = "Hello"
}

resource "aws_api_gateway_resource" "hello" {
  path_part   = "hello"
  parent_id   = "${aws_api_gateway_rest_api.hello.root_resource_id}"
  rest_api_id = "${aws_api_gateway_rest_api.hello.id}"
}

resource "aws_api_gateway_method" "hello" {
  rest_api_id   = "${aws_api_gateway_rest_api.hello.id}"
  resource_id   = "${aws_api_gateway_resource.hello.id}"
  http_method   = "POST"
  authorization = "NONE"
}

data "aws_region" "current" {}

resource "aws_api_gateway_integration" "hello" {
  rest_api_id = "${aws_api_gateway_rest_api.hello.id}"
  resource_id = "${aws_api_gateway_resource.hello.id}"
  http_method = "${aws_api_gateway_method.hello.http_method}"

  # Lambda can only do POST
  integration_http_method = "POST"

  type = "AWS_PROXY"
  uri  = "arn:aws:apigateway:${data.aws_region.current.name}:lambda:path/2015-03-31/functions/${aws_lambda_function.hello.arn}/invocations"
}

data "aws_caller_identity" "current" {}

resource "aws_lambda_permission" "hello" {
  statement_id = "AllowExecutionFromAPIGateway"
  action       = "lambda:InvokeFunction"
  principal    = "apigateway.amazonaws.com"

  function_name = "${aws_lambda_function.hello.arn}"

  source_arn = "arn:aws:execute-api:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:${aws_api_gateway_rest_api.hello.id}/*/${aws_api_gateway_method.hello.http_method}${aws_api_gateway_resource.hello.path}"
}

resource "aws_api_gateway_deployment" "hello" {
  rest_api_id = "${aws_api_gateway_rest_api.hello.id}"
  stage_name  = "test"

  # This is necessary to avoid a race condition
  depends_on = ["aws_api_gateway_integration.hello"]
}

output "invoke_url" {
  value = "${aws_api_gateway_deployment.hello.invoke_url}"
}
