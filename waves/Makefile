PROTOFILES=../protobuf/waveform.proto
GOPROTOFILES=$(patsubst %.proto,%.pb.go,$(PROTOFILES))

CLEANFILES+=$(GOPROTOFILES)

protobuf: $(PROTOFILES)
	protoc -I=../protobuf --go_out=./internal/protobuf --go_opt=paths=source_relative --go-grpc_out=./internal/protobuf --go-grpc_opt=paths=source_relative $<

.PHONY: protobuf