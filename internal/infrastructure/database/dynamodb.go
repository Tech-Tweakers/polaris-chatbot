package database

import (
	"context"
	"fmt"
	"sort"

	"ecatrom/internal/domain/ecatrom"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func NewDynamoDB(cfg aws.Config, tableName string) ecatrom.Repository {
	return &dynamoDB{
		db:        dynamodb.NewFromConfig(cfg),
		tableName: tableName,
	}
}

type dynamoDB struct {
	tableName string
	db        *dynamodb.Client
}

func (d *dynamoDB) Find(key string) (*ecatrom.ChatPersistence, error) {
	result, err := d.db.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(d.tableName),
		Key: map[string]types.AttributeValue{
			"key": &types.AttributeValueMemberS{Value: key},
		},
	})

	fmt.Printf("result: %v\n", result)

	if err != nil {
		return nil, nil
	}

	if result != nil && result.Item == nil {
		return nil, fmt.Errorf("entry not found in database")
	}

	dbItem := ecatrom.ChatPersistence{}

	err = attributevalue.UnmarshalMap(result.Item, &dbItem)

	if err != nil {
		return nil, err
	}
	return &dbItem, nil
}

func (d *dynamoDB) Insert(applicationEntity ecatrom.ChatPersistence) (*ecatrom.ChatPersistence, error) {
	return d.upsert(applicationEntity)
}

func (d *dynamoDB) Upsert(applicationEntity ecatrom.ChatPersistence) (*ecatrom.ChatPersistence, error) {
	return d.upsert(applicationEntity)
}

func (d *dynamoDB) Delete(key string) error {
	_, err := d.db.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String(d.tableName),
		Key: map[string]types.AttributeValue{
			"key": &types.AttributeValueMemberS{Value: key},
		},
	})

	return err
}

func (d *dynamoDB) upsert(ChatPersistence ecatrom.ChatPersistence) (*ecatrom.ChatPersistence, error) {

	item, err := attributevalue.MarshalMap(ChatPersistence)
	if err != nil {
		fmt.Println("error when try to convert user data to dynamodbattribute")
		return nil, err
	}
	item["key"] = &types.AttributeValueMemberS{Value: fmt.Sprint(ChatPersistence.EntryID)}

	params := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(d.tableName),
	}

	if _, err := d.db.PutItem(context.Background(), params); err != nil {
		fmt.Println("error r.db.PutItem")
		return nil, err
	}
	return &ChatPersistence, nil
}

func (d *dynamoDB) List() (*[]ecatrom.ChatPersistence, error) {
	result, err := d.db.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String(d.tableName),
	})

	if err != nil {
		return nil, err
	}

	var records []ecatrom.ChatPersistence

	err = attributevalue.UnmarshalListOfMaps(result.Items, &records)

	if err != nil {
		return nil, err
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].EntryID < records[j].EntryID
	})

	return &records, nil
}
