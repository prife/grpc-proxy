package main

import (
	"context"
	"fmt"

	"github.com/mwitkow/grpc-proxy/perfdog/perfdog"
)

// defaultPerfDogServer is the canonical implementation of a TestServiceServer.
type defaultPerfDogServer struct {
	perfdog.UnsafePerfDogServiceServer
	// fakeperfdog.UnimplementedPerfDogServiceServer
}

var defaultPerfDogServiceServer = defaultPerfDogServer{}

var perfCC perfdog.PerfDogServiceClient

func (s defaultPerfDogServer) LoginWithToken(ctx context.Context, token *perfdog.Token) (*perfdog.UserInfo, error) {
	fmt.Printf("LoginWithToken: %s\n", token.Token)
	if perfCC == nil {
		return nil, fmt.Errorf("perfCC is nil")
	}
	return perfCC.LoginWithToken(ctx, token)
}
