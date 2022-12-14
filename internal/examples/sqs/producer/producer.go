package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/toretto460/squeue"
	sqsexample "github.com/toretto460/squeue/internal/examples/sqs"
	"github.com/toretto460/squeue/sqs"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if !sqsexample.CheckEnvVariables("AWS_PROFILE", "AWS_SHARED_CREDENTIALS_FILE", "AWS_REGION", "AWS_QUEUE_URL") {
		log.Fatal(`Please set the env variables
		AWS_PROFILE=user-dev-admin
		AWS_SHARED_CREDENTIALS_FILE=/Users/{name.lastname}/.aws/credentials
		AWS_REGION=eu-central-1
		AWS_QUEUE_URL=https://sqs.eu-central-1.amazonaws.com/...
		`)
	}
}

func main() {
	d, err := sqs.New(
		sqs.WithUrl(os.Getenv("AWS_QUEUE_URL")),
		sqs.WithRegion(os.Getenv("AWS_REGION")),
		sqs.WithMaxNumberOfMessages(10_000),
	)

	if err != nil {
		panic(err)
	}

	pub := squeue.NewQueue(d, "test-simone")

	pub.Enqueue(&sqsexample.MyEvent{Name: "foo"})
	pub.Enqueue(&sqsexample.MyEvent{Name: "bar"})
	pub.Enqueue(&sqsexample.MyEvent{Name: "baz"})
}
