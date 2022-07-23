package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	DefaultProbeDuration = 5 * time.Second
)

func Probe(opts *ClientOptions, timeout time.Duration) error {
	client, err := initializeMongoClient(opts)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	err = client.Ping(ctx, readpref.Primary())
	defer cancel()
	return err
}
