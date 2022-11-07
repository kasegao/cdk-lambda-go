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

## Useful commands

- `cdk deploy` deploy this stack to your default AWS account/region
- `cdk diff` compare deployed stack with current state
- `cdk synth` emits the synthesized CloudFormation template
- `go test` run unit tests
