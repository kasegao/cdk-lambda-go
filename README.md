# cdk-lambda-go

Sample CDK script with **Go** to create a Lambda function in **Go**.

## Structure

```bash
.
├── cdk-lambda-go.go       # main code to define stacks
├── cdk-lambda-go_test.go  # test code
├── cdk.json               # CDK settings
└── lambda                 # resources for Lambda
    └── handler.go
```

## Quick Start

Deploy Stack

```bash
cdk deploy
```

Test

```bash
aws sqs send-message --queue-url <SQS URL> --message-body "Test message."
```

then, you can find the following log in CloudWatch:

```txt
The message <MessageId> for event source aws:sqs = Test message.
```

## Useful commands

- `cdk deploy` deploy this stack to your default AWS account/region
- `cdk diff` compare deployed stack with current state
- `cdk synth` emits the synthesized CloudFormation template
- `go test` run unit tests
