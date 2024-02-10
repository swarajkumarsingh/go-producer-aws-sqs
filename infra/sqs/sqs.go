package sqs

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/swarajkumarsingh/go-producer-aws-sqs/conf"
)

var sess, _ = session.NewSessionWithOptions(session.Options{
	Config: aws.Config{
		Region:      aws.String(conf.AWS_REGION),
		Credentials: credentials.NewStaticCredentials(conf.AWS_ACCESS_KEY, conf.AWS_SECRET_ACCESS_KEY, conf.AWS_TOKEN),
	},
})

var svc = sqs.New(sess)

// SendFIFOMessage sends events to SQS queue
func SendFIFOMessage(messageID string, messageBody string, messageGroupID string, queueURL string) error {
	_, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds:           aws.Int64(0),
		MessageBody:            &messageBody,
		MessageGroupId:         &messageGroupID,
		MessageDeduplicationId: &messageID,
		QueueUrl:               &queueURL,
	})
	return err
}
