package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Connect to AWS resources using role and publish topic
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		profileConfig = os.Getenv("PROFILE_CONFIG")
		targetRole    = os.Getenv("TARGET_ROLE")
		externalId    = os.Getenv("EXTERNAL_ID")
		snsTopicArn   = os.Getenv("SNS_TOPIC_ARN")
	)

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(profileConfig))

	if err != nil {
		fmt.Println("error", err)
		return
	}

	stsSvc := sts.NewFromConfig(cfg)
	creds := stscreds.NewAssumeRoleProvider(stsSvc, targetRole, func(o *stscreds.AssumeRoleOptions) {
		o.ExternalID = aws.String(externalId)
	})

	cfg.Credentials = aws.NewCredentialsCache(creds)

	// SNS
	topicPtr := snsTopicArn
	msgPtr := "Hii"

	svc := sns.NewFromConfig(cfg)

	result, err := svc.Publish(context.TODO(), &sns.PublishInput{
		Message:  &msgPtr,
		TopicArn: &topicPtr,
	})

	if err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Println(*result.MessageId)

	// Get topic attributes
	//allTopics, err := svc.GetTopicAttributes(context.TODO(), &sns.GetTopicAttributesInput{TopicArn: &topicPtr})
	//
	//if err != nil {
	//	fmt.Println("error", err)
	//	return
	//}
	//
	//fmt.Println(allTopics.Attributes)
}
