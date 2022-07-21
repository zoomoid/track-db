# track-db/waveman

waveman for track-db is a gRPC server that takes a file reference to a shared MongoDB grid FS and
configuration properties matching those of waveman2, and renders
such a waveform as SVG using <https://github.com/zoomoid/waveman2>.

The message format is defined as protobuf and mapped to structs that waveman understands.