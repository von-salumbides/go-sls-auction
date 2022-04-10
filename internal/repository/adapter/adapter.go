package adapter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/von-salumbides/auction/utils/logger"
)

type Database struct {
	connection *dynamodb.DynamoDB
}

type Interface interface {
	Health(tableName string) (bool, error)
	CreateOrUpdate(entity interface{}, tableName string) (*dynamodb.PutItemOutput, error)
	GetAll(tableName string) (*dynamodb.ScanOutput, error)
	Get(pathParam string, tableName string) (*dynamodb.GetItemOutput, error)
}

// NewAdapter creates new Dynamodb adapter
func NewAdapter(con *dynamodb.DynamoDB) Interface {
	return &Database{
		connection: con,
	}
}

// Health check for Dynamodb
func (db *Database) Health(tableName string) (bool, error) {
	_, err := db.connection.ListTables(&dynamodb.ListTablesInput{
		ExclusiveStartTableName: &tableName,
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

// CreateOrUpdate PUT data to dynamodb
func (db *Database) CreateOrUpdate(entity interface{}, tableName string) (*dynamodb.PutItemOutput, error) {
	// entity parsed
	entityParsed, err := dynamodbattribute.MarshalMap(entity)
	if err != nil {
		logger.ERROR("Marshal failure", err.Error())
		return nil, err
	}
	// Build input item
	input := &dynamodb.PutItemInput{
		Item:      entityParsed,
		TableName: aws.String(tableName),
	}
	return db.connection.PutItem(input)
}

func (db *Database) GetAll(tableName string) (*dynamodb.ScanOutput, error) {
	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}
	return db.connection.Scan(params)
}

func (db *Database) Get(pathParam string, tableName string) (*dynamodb.GetItemOutput, error) {
	input := &dynamodb.GetItemInput{
		TableName: &tableName,
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(pathParam),
			},
		},
	}
	request, err := db.connection.GetItem(input)
	if err != nil {
		logger.ERROR("Get item failed", err.Error())
	}
	return request, nil
}
