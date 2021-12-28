package dynamo

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/shinnosuke-K/github-dev-insight/pkg/infrastructure/aws/dynamo/client"
)

type Config struct {
	SecretAccessKey string
	AccessToken     string
	Region          string
	EndPoint        string
}

func NewDynamoDB(cfg Config) (*client.Client, error) {
	switch {
	case cfg.SecretAccessKey == "":
		return nil, errors.New("secrete access key is required")
	case cfg.AccessToken == "":
		return nil, errors.New("access token is required")
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(cfg.Region),
		Credentials: credentials.NewStaticCredentials(cfg.AccessToken, cfg.SecretAccessKey, ""),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create session. %w", err)
	}

	dynamo := dynamodb.New(sess, &aws.Config{
		Endpoint: aws.String(cfg.EndPoint),
		Region:   aws.String(cfg.Region),
	})
	return &client.Client{DB: dynamo}, nil
}
