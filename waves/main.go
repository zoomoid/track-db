package main

import (
	"log"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/zoomoid/track-db/waves/internal/grpc"
	"github.com/zoomoid/track-db/waves/internal/mongodb"
)

func init() {
	viper.SetEnvPrefix("waves")
	viper.AutomaticEnv()
	viper.BindEnv("port")
	viper.BindEnv("mongo-uri", "WAVES_MONGO_URI")
	viper.BindEnv("mongo-username", "WAVES_MONGO_USERNAME")
	viper.BindEnv("mongo-password", "WAVES_MONGO_PASSWORD")
	viper.BindEnv("mongo-database", "WAVES_MONGO_DATABASE")
	viper.BindEnv("mongo-collection", "WAVES_MONGO_COLLECTION")

	pflag.Uint16("port", 9000, "Port to run the gRPC server on")
	pflag.String("mongo-uri", "mongodb://localhost:27017/", "MongoDB connection string URI")
	pflag.String("mongo-username", "", "MongoDB username")
	pflag.String("mongo-password", "", "MongoDB password")
	pflag.String("mongo-database", "track-db", "MongoDB database to connect to")
	pflag.String("mongo-collection", "fs.files", "GridFS collection name to use")
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)
}

func main() {
	mongoCredential := &mongodb.ClientOptions{
		URI:      viper.GetString("mongo-uri"),
		Username: viper.GetString("mongo-username"),
		Password: viper.GetString("mongo-password"),
	}

	// before starting any server, probe mongodb for availability, otherwise exit
	err := mongodb.Probe(mongoCredential, mongodb.DefaultProbeDuration)

	if err != nil {
		log.Fatalf("failed to connect to mongodb: %v", err)
	}

	grpc.StartGRPCServer(uint16(viper.GetUint32("port")))
}
