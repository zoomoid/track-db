package grpc

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/spf13/viper"
	"github.com/zoomoid/track-db/waves/internal/mongodb"
	pb "github.com/zoomoid/track-db/waves/internal/protobuf"
	"github.com/zoomoid/track-db/waves/internal/runner"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedWavemanServer
}

func (s *server) Waveform(ctx context.Context, in *pb.TrackRequest) (*pb.TrackResponse, error) {
	ro := ConvertProtobufToWavemanSpec(in)

	gridFSCFG := &mongodb.GridFSClientConfiguration{
		URI:      viper.GetString("mongo_uri"),
		Username: viper.GetString("mongo_username"),
		Password: viper.GetString("mongo_password"),
	}

	gfsClient, err := mongodb.NewGridFSClient(gridFSCFG)
	if err != nil {
		return nil, err
	}

	buf, err := gfsClient.
		DB(viper.GetString("mongo_database")).
		Collection(viper.GetString("mongo_collection")).
		DownloadFile(in.GetId())

	if err != nil {
		return nil, err
	}

	ro.File = buf

	runner := runner.NewRunner(ro)

	runner.Transform().Paint()
	if err := runner.Error(); err != nil {
		return nil, err
	}

	svg := runner.SVG().String()
	samples := runner.Samples().Bytes()

	return &pb.TrackResponse{
		Id:      in.GetId(),
		Svg:     svg,
		Samples: samples,
	}, nil
}

func StartGRPCServer(port uint16) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterWavemanServer(s, &server{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
