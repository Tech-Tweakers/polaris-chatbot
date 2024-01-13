package database

import (
	"context"
	"sort"

	"polarisai/internal/domain/polaris"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDB(ctx context.Context, connectionString, dbName, collectionName string) (polaris.Repository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	return &mongoDB{
		client:         client,
		dbName:         dbName,
		collectionName: collectionName,
	}, nil
}

type mongoDB struct {
	client         *mongo.Client
	dbName         string
	collectionName string
}

func (m *mongoDB) Insert(applicationEntity polaris.ChatPersistence) (*polaris.ChatPersistence, error) {
	return m.upsert(applicationEntity)
}

func (m *mongoDB) Upsert(applicationEntity polaris.ChatPersistence) (*polaris.ChatPersistence, error) {
	return m.upsert(applicationEntity)
}

func (m *mongoDB) upsert(ChatPersistence polaris.ChatPersistence) (*polaris.ChatPersistence, error) {
	collection := m.client.Database(m.dbName).Collection(m.collectionName)

	filter := bson.D{{Key: "entryid", Value: ChatPersistence.EntryID}}
	update := bson.D{{Key: "$set", Value: ChatPersistence}}

	_, err := collection.UpdateOne(context.TODO(), filter, update, options.Update().SetUpsert(true))
	if err != nil {
		return nil, err
	}

	return &ChatPersistence, nil
}

func (m *mongoDB) List() (*[]polaris.ChatPersistence, error) {
	collection := m.client.Database(m.dbName).Collection(m.collectionName)
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var records []polaris.ChatPersistence
	err = cursor.All(context.TODO(), &records)
	if err != nil {
		return nil, err
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].EntryID < records[j].EntryID
	})

	return &records, nil
}
