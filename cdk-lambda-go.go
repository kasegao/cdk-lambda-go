package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdaeventsources"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type CdkLambdaGoStackProps struct {
	awscdk.StackProps
}

func NewCdkLambdaGoStack(scope constructs.Construct, id string, props *CdkLambdaGoStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	queue := awssqs.NewQueue(stack, jsii.String("Queue"), &awssqs.QueueProps{
		VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(60)),
	})

	fn := awslambda.NewFunction(stack, jsii.String("Lambda"), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_GO_1_X(),
		Code: awslambda.AssetCode_FromAsset(jsii.String("lambda"), &awss3assets.AssetOptions{
			Bundling: &awscdk.BundlingOptions{
				Image:   awslambda.Runtime_GO_1_X().BundlingImage(),
				Command: jsii.Strings("bash", "-c", "GOOS=linux GOARCH=amd64 go build -o handler"),
				Local:   &LocalBundling{},
			},
		}),
		Handler: jsii.String("handler"),
		Timeout: awscdk.Duration_Seconds(jsii.Number(30)),
	})
	fn.AddEventSource(awslambdaeventsources.NewSqsEventSource(queue, &awslambdaeventsources.SqsEventSourceProps{
		BatchSize: jsii.Number(10),
	}))

	return stack
}

// Local build settings for lambda function with golang
type LocalBundling struct{}

func (*LocalBundling) TryBundle(outputDir *string, options *awscdk.BundlingOptions) *bool {
	path := filepath.Join(*outputDir, "handler")
	var err error
	if runtime.GOOS == "windows" {
		err = exec.Command("cmd", "/c", fmt.Sprintf("set GOOS=linux&set GOARCH=amd64&cd lambda&go build -o %s", path)).Run()
	} else {
		err = exec.Command("bash", "-c", fmt.Sprintf("GOOS=linux GOARCH=amd64 cd lambda && go build -o %s", path)).Run()
	}
	return jsii.Bool(err == nil)
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewCdkLambdaGoStack(app, "CdkLambdaGoStack", &CdkLambdaGoStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String("012345678901"),
		Region:  jsii.String("ap-northeast-1"),
	}
}
