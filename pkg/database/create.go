package database

import (
	"context"
	"fmt"
	logs "stream/pkg/utils/logging"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go/aws"
)

// Gets tables names from dynamo
func (dynamoX *MyDynamoClient) ListTables() ([]string, error) {

	resp, err := dynamoX.client.ListTables(context.TODO(), &dynamodb.ListTablesInput{
		Limit: aws.Int32(5),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list tables, %v", err)
	}

	return resp.TableNames, nil
}

// Adds a entry to dynamo
func (dynamoX *MyDynamoClient) AddEntry() error {
	// Prepare the input for the PutItem operation
	input := &dynamodb.PutItemInput{
		TableName: aws.String(dynamoX.TableName),
		Item: map[string]types.AttributeValue{
			"videoId": &types.AttributeValueMemberS{Value: "test"},
			"Value":   &types.AttributeValueMemberS{Value: "axolot"},
		},
	}

	// Execute the PutItem operation
	_, err := dynamoX.client.PutItem(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("failed to add entry on dynamo, %v", err)
	}

	logs.I.Println("Entry added successfully.")
	return nil
}

// Reads the content of an entry from dynamo
func (dynamoX *MyDynamoClient) ReadEntry(PartitionKey string) error {
	// Prepare the input for the GetItem operation
	input := &dynamodb.GetItemInput{
		TableName: aws.String(dynamoX.TableName),
		Key: map[string]types.AttributeValue{
			"videoId": &types.AttributeValueMemberS{Value: PartitionKey},
		},
	}

	// Execute the GetItem operation
	result, err := dynamoX.client.GetItem(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("failed to read entry, %v", err)
	}

	if valueAttr, ok := result.Item["Value"].(*types.AttributeValueMemberS); ok {
		//SE HACE UNA ASertion
		//todo Aprender a hacer eso
		// Now valueAttr.Value contains the clean string
		logs.I.Println("Entry:", valueAttr.Value)
		return nil
	} else {
		// Handle the case where the value is not of the expected type or is missing
		return fmt.Errorf("Entry value is missing or not a string")

	}

}
