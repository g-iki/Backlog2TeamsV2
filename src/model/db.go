package db

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

var (
	gdb    *dynamo.DB
	region string

	usersTable string
)

var DB *dynamo.DB

func init() {

	dynamoDbRegion := "ap-northeast-1"
	disableSsl := false

	// DynamoDB Localを利用する場合はEndpointのURLを設定する
	dynamoDbEndpoint := os.Getenv("DYNAMO_ENDPOINT")
	if len(dynamoDbEndpoint) != 0 {
		disableSsl = true
	}

	DB = dynamo.New(
		session.New(),
		&aws.Config{
			Region:     aws.String(dynamoDbRegion),
			Endpoint:   aws.String(dynamoDbEndpoint),
			DisableSSL: aws.Bool(disableSsl),
		},
	)

}
