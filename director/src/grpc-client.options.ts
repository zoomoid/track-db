import { ClientOptions, Transport } from "@nestjs/microservices";
import { join } from "path";

export const grpcClientOptions: ClientOptions = {
  transport: Transport.GRPC,
  options: {
    package: "waves",
    protoPath: join(__dirname, "./waves/track.proto")
  }
}