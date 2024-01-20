package main

import (
  hook "go-emqx-ex-hook-to-tdengine/hook"
  pb "go-emqx-ex-hook-to-tdengine/protobuf"
  "google.golang.org/grpc"
  "log"
  "net"
)

const (
  port = ":9000"
)

func main() {
  lis, err := net.Listen("tcp", port)
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  s := grpc.NewServer()
  pb.RegisterHookProviderServer(s, &hook.Server{})
  log.Println("Started gRPC server on ::9000")
  if err := s.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}
