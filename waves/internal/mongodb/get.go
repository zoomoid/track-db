package mongodb

import (
	"bytes"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

type GridFSClientConfiguration struct {
	ClientOptions
	db         string
	collection string
}

type GridFSClient struct {
	c          *GridFSClientConfiguration
	db         *mongo.Database
	collection *mongo.Collection
	conn       *mongo.Client
	bucket     *gridfs.Bucket
}

func NewGridFSClient(cfg *ClientOptions) (*GridFSClient, error) {
	c := &GridFSClient{
		c: &GridFSClientConfiguration{},
	}
	c.c.ClientOptions = *cfg
	client, err := c.client()
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

func (g *GridFSClient) client() (*mongo.Client, error) {
	return initializeMongoClient(&ClientOptions{
		URI:      g.c.URI,
		Username: g.c.Username,
		Password: g.c.Password,
	})
}
