package main

import (
    "gopher-school/lesson5-grpc-adder/pkg/adder"
    "gopher-school/lesson5-grpc-adder/pkg/api"
    "google.golang.org/grpc"
    "net"
    "log"
)

func main() {
    s := grpc.NewServer()
    srv := &adder.GRPCServer{}
    api.RegisterAdderServer(s, srv)

    l, err := net.Listen("tcp", ":4848")
    if err != nil {
        log.Fatal(err)
    }

    if err := s.Serve(l); err != nil {
        log.Fatal(err)
    }
}
