package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func main() {

	my_table := "my-table"
	id := "679374c6-d940-4cb5-9422-9a7544487ac7"

	// Using the SDK's default configuration, loading additional config
	// and credentials values from the environment variables, shared
	// credentials, and shared configuration files
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("sa-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Using the Config value, create the DynamoDB client
	client := dynamodb.NewFromConfig(cfg)

	// Define the input parameters
	input := &dynamodb.QueryInput{
		TableName:              aws.String(my_table),
		KeyConditionExpression: aws.String("id = :id"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":id": &types.AttributeValueMemberS{Value: id},
		},
		ScanIndexForward: aws.Bool(false),
		// IndexName:              aws.String("datetime"),
		Limit:                  aws.Int32(1),
		ConsistentRead:         aws.Bool(false),
		ReturnConsumedCapacity: types.ReturnConsumedCapacityNone,
	}

	// Make the Query API call to DynamoDB
	result, err := client.Query(context.TODO(), input)
	if err != nil {
		fmt.Println("Error querying DynamoDB table:", err)
		return
	}

	// Check if any items were returned
	if len(result.Items) == 0 {
		fmt.Println("No items found for the given ID")
		return
	}

	// Get the last item
	lastItem := result.Items[0]

	// Print the item details
	fmt.Println("ID:", lastItem["id"])
	fmt.Println("Status:", lastItem["status"])
	fmt.Println("Datetime:", lastItem["datetime"])
}
