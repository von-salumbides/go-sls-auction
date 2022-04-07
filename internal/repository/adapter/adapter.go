package adapter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"go.uber.org/zap"
)

type Database struct {
	connection *dynamodb.DynamoDB
}

type Interface interface {
	Health() bool
	CreateOrUpdate(entity interface{}, tableName string) (*dynamodb.PutItemOutput, error)
}

func NewAdapter(con *dynamodb.DynamoDB) Interface {
	return &Database{
		connection: con,
	}
}

func (db *Database) Health() bool {
	_, err := db.connection.ListTables(&dynamodb.ListTablesInput{})
	return err == nil
}

func (db *Database) CreateOrUpdate(entity interface{}, tableName string) (*dynamodb.PutItemOutput, error) {
	// entity parsed
	entityParsed, err := dynamodbattribute.MarshalMap(entity)
	if err != nil {
		zap.L().Fatal("Marshal failure", zap.Any("error", err.Error()))
		return nil, err
	}

	// Build input item
	input := &dynamodb.PutItemInput{
		Item:      entityParsed,
		TableName: aws.String(tableName),
	}
	return db.connection.PutItem(input)
}
