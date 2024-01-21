package aws

import (
	"context"
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

type SQSReceiveMessageAPI interface {
	GetQueueUrl(ctx context.Context,
		params *sqs.GetQueueUrlInput,
		optFns ...func(*sqs.Options)) (*sqs.GetQueueUrlOutput, error)

	ReceiveMessage(ctx context.Context,
		params *sqs.ReceiveMessageInput,
		optFns ...func(*sqs.Options)) (*sqs.ReceiveMessageOutput, error)

	DeleteMessage(ctx context.Context,
		params *sqs.DeleteMessageInput,
		optFns ...func(*sqs.Options)) (*sqs.DeleteMessageOutput, error)
}

func GetQueueURL(c context.Context, api SQSReceiveMessageAPI, input *sqs.GetQueueUrlInput) (*sqs.GetQueueUrlOutput, error) {
	return api.GetQueueUrl(c, input)
}

func GetMessages(c context.Context, api SQSReceiveMessageAPI, input *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
	return api.ReceiveMessage(c, input)
}

func DeleteMessage(c context.Context, api SQSReceiveMessageAPI, input *sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error) {
	return api.DeleteMessage(c, input)
}

func ReceiveMessage() {
	println("start ReceiveMessage")
	queue := flag.String("q", "riadq", "The name of the queue")
	timeout := flag.Int("t", 5, "How long, in seconds, that the message is hidden from others")
	flag.Parse()

	if *queue == "" {
		fmt.Println("You must supply the name of a queue (-q QUEUE)")
		return
	}

	if *timeout < 0 {
		*timeout = 0
	}

	if *timeout > 12*60*60 {
		*timeout = 12 * 60 * 60
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := sqs.NewFromConfig(cfg)

	gQInput := &sqs.GetQueueUrlInput{
		QueueName: queue,
	}

	// Get URL of queue
	urlResult, err := GetQueueURL(context.TODO(), client, gQInput)
	if err != nil {
		fmt.Println("Got an error getting the queue URL:")
		fmt.Println(err)
		return
	}

	queueURL := urlResult.QueueUrl

	gMInput := &sqs.ReceiveMessageInput{
		MessageAttributeNames: []string{
			string(types.QueueAttributeNameAll),
		},
		QueueUrl:            queueURL,
		MaxNumberOfMessages: 1,
		VisibilityTimeout:   int32(*timeout),
		WaitTimeSeconds:     int32(3),
	}

	for {

		msgResult, err := GetMessages(context.TODO(), client, gMInput)
		if err != nil {
			fmt.Println("Got an error receiving messages:")
			fmt.Println(err)
			return
		}

		if msgResult.Messages != nil {
			fmt.Println("Message ID:     " + *msgResult.Messages[0].MessageId)
			fmt.Println("Message Handle: " + *msgResult.Messages[0].ReceiptHandle)
			fmt.Println("Message Body:   " + *msgResult.Messages[0].Body)
		} else {
			fmt.Println("No messages found")
			fmt.Println("end ReceiveMessage")
		}

		if msgResult.Messages == nil {
			continue
		}

		dMInput := &sqs.DeleteMessageInput{
			QueueUrl:      queueURL,
			ReceiptHandle: msgResult.Messages[0].ReceiptHandle,
		}

		_, err = DeleteMessage(context.TODO(), client, dMInput)
		if err != nil {
			fmt.Println("Got an error deleting the message:")
			fmt.Println(err)
			return
		}

		fmt.Println("end ReceiveMessage")
	}
}
