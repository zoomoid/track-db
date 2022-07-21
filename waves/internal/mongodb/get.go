package mongodb

import (
	"bytes"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GridFSClientConfiguration struct {
	URI        string
	db         string
	collection string
	Username   string
	Password   string
}

type GridFSClient struct {
	c          *GridFSClientConfiguration
	db         *mongo.Database
	collection *mongo.Collection
	conn       *mongo.Client
	bucket     *gridfs.Bucket
}

func NewGridFSClient(cfg *GridFSClientConfiguration) (*GridFSClient, error) {
	c := &GridFSClient{
		c: cfg,
	}
	client, err := c.initializeMongoClient()
	if err != nil {
		return nil, err
	}
	c.conn = client
	return c, nil
}

func (g *GridFSClient) DB(db string) *GridFSClient {
	g.c.db = db
	g.db = g.conn.Database(g.c.db)
	return g
}

func (g *GridFSClient) Collection(collection string) *GridFSClient {
	g.c.collection = collection
	g.collection = g.db.Collection(g.c.collection)
	return g
}

func (g *GridFSClient) DownloadFile(id string) (*bytes.Buffer, error) {
	// if bucket is not yet set, set it
	if g.bucket == nil {
		bucket, _ := gridfs.NewBucket(g.db)
		g.bucket = bucket
	}

	var buf bytes.Buffer
	// TODO: check if this needs bson.ObjectID conversion before
	_, err := g.bucket.DownloadToStream(id, &buf)
	if err != nil {
		return nil, err
	}

	return &buf, nil
}

func (g *GridFSClient) initializeMongoClient() (*mongo.Client, error) {
	var err error
	var client *mongo.Client

	credentials := options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		Username:      g.c.Username,
		Password:      g.c.Password,
	}

	opts := options.Client().
		ApplyURI(g.c.URI).
		SetMaxPoolSize(5).
		SetAuth(credentials)

	if client, err = mongo.Connect(context.Background(), opts); err != nil {
		return nil, err
	}
	return client, nil
}
