package dynamo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Client struct {
    Service *dynamodb.Client
}

func New() (*Client, error) {
    cfg, err := config.LoadDefaultConfig(context.TODO())
    if err != nil {
        return nil, err
    }
    return &Client{Service: dynamodb.NewFromConfig(cfg)}, nil
}
