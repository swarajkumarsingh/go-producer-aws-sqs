package main

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/swarajkumarsingh/go-producer-aws-sqs/conf"
	"github.com/swarajkumarsingh/go-producer-aws-sqs/infra/sqs"
)

func main() {
	eventID := uuid.New().String()
	message := "swaraj.singh.wearingo@gmail.com"

	timer := time.Now()
	err := sqs.SendFIFOMessage(eventID, message, conf.SQSMessageGroupBilling, conf.AWS_SQS_URL)
	if err != nil {
		panic(err)
	}

	log.Println("message successfully added to the queue and took", time.Since(timer))
}
