package main

import (
	"context"
	"fmt"
	"net"
	"time"

	// "github.com/mwitkow/grpc-proxy/perfdog/fakeperfdog"
	"github.com/mwitkow/grpc-proxy/perfdog/perfdog"
	"github.com/mwitkow/grpc-proxy/proxy"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	backendCC, err := grpc.DialContext(ctx,
		"localhost:23456", // perfdog addr
		grpc.WithInsecure(),
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(10*1024*1024)))
	if err != nil {
		panic(err)
		//return nil, fmt.Errorf("dialing backend: %v", err)
	}

	perfCC = perfdog.NewPerfDogServiceClient(backendCC)
	// First, we need to create a client connection to this backend.
	// proxySrv := proxy.NewProxy(backendCC)

	directorFn := func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
		fmt.Printf("--> fullMethodName: %s\n", fullMethodName)
		md, _ := metadata.FromIncomingContext(ctx)
		fmt.Printf("--> fullMethodName: %s, md=%#v\n", fullMethodName, md)
		outCtx := metadata.NewOutgoingContext(ctx, md.Copy())
		return outCtx, backendCC, nil
	}
	proxySrv := grpc.NewServer(grpc.UnknownServiceHandler(proxy.TransparentHandler(directorFn)))
	perfdog.RegisterPerfDogServiceServerMini(proxySrv, defaultPerfDogServiceServer)
	proxyBc, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9002))
	if err != nil {
		panic(err)
	}

	// run the proxy backend
	if err := proxySrv.Serve(proxyBc); err != nil {
		if err == grpc.ErrServerStopped {
			return
		}
	}
	proxySrv.GracefulStop()
}
