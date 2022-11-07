package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	MessageId   string `json:"messageId"`
	EventSource string `json:"eventSource"`
	Body        string `json:"body"`
}

func (r *Response) logger() string {
	return fmt.Sprintf("The message %s for event source %s = %s \n", r.MessageId, r.EventSource, r.Body)
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) ([]*Response, error) {
	var responses []*Response
	for _, message := range sqsEvent.Records {
		resp := &Response{
			MessageId:   message.MessageId,
			EventSource: message.EventSource,
			Body:        message.Body,
		}
		responses = append(responses, resp)
		fmt.Println(resp.logger())
	}

	return responses, nil
}

func main() {
	lambda.Start(handler)
}
