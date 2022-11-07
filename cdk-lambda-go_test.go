package main

import (
	"os"
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
)

func TestCdkLambdaGoStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack := NewCdkLambdaGoStack(app, "MyStack", nil)

	// THEN
	template := assertions.Template_FromStack(stack, nil)

	template.HasResourceProperties(jsii.String("AWS::SQS::Queue"), map[string]interface{}{
		"VisibilityTimeout": 60,
	})

	template.HasResourceProperties(jsii.String("AWS::Lambda::Function"), map[string]interface{}{
		"Handler": "handler",
		"Runtime": "go1.x",
		"Timeout": 30,
	})

	template.HasResourceProperties(jsii.String("AWS::Lambda::EventSourceMapping"), map[string]interface{}{
		"BatchSize": 10,
	})
}

func TestLambdaLocalBuild(t *testing.T) {
	d := t.TempDir()
	t.Logf("outputDir: %s", d)

	b := &LocalBundling{}
	b.TryBundle(&d, nil)

	files, err := os.ReadDir(d)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 1 {
		t.Fatalf("expected 1 file, got %d", len(files))
	}
	if files[0].Name() != "handler" {
		t.Fatalf("expected handler, got %s", files[0].Name())
	}
}
