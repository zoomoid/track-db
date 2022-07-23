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

	viper.BindEnv("port")
	viper.BindEnv("mongo_uri")
	viper.BindEnv("mongo_username")
	viper.BindEnv("mongo_password")
	viper.BindEnv("mongo_database")
	viper.BindEnv("mongo_collection")

	pflag.Uint16("port", 9000, "Port to run the gRPC server on")
	pflag.String("mongo_uri", "mongodb://localhost:27017/", "MongoDB connection string URI")
	pflag.String("mongo_username", "", "MongoDB username")
	pflag.String("mongo_password", "", "MongoDB password")
	pflag.String("mongo_database", "track-db", "MongoDB database to connect to")
	pflag.String("mongo_collection", "fs.files", "GridFS collection name to use")

	viper.BindPFlags(pflag.CommandLine)
}

func main() {
	// before starting any server, probe mongodb for availability, otherwise exit
	err := mongodb.Probe(&mongodb.ClientOptions{
		URI:      viper.GetString("mongo_uri"),
		Username: viper.GetString("mongo_username"),
		Password: viper.GetString("mongo_password"),
	}, mongodb.DefaultProbeDuration)
	if err != nil {
		log.Fatalf("failed to connect to mongodb: %v", err)
	}

	grpc.StartGRPCServer(uint16(viper.GetUint32("port")))
}
