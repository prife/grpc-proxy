package main

import (
	"context"
	"fmt"

	"github.com/mwitkow/grpc-proxy/perfdog/fakeperfdog"
	"github.com/mwitkow/grpc-proxy/perfdog/perfdog"
)

// defaultPerfDogServer is the canonical implementation of a TestServiceServer.
type defaultPerfDogServer struct {
	// perfdog.UnsafePerfDogServiceServer
	fakeperfdog.UnimplementedPerfDogServiceServer
}

var defaultPerfDogServiceServer = defaultPerfDogServer{}

var perfCC perfdog.PerfDogServiceClient

func (s defaultPerfDogServer) LoginWithToken(ctx context.Context, token *fakeperfdog.Token) (*fakeperfdog.UserInfo, error) {
	fmt.Printf("LoginWithToken: %s\n", token.Token)
	if perfCC == nil {
		return nil, fmt.Errorf("perfCC is nil")
	}
	perfCC.LoginWithToken(ctx, &perfdog.Token{Token: token.Token})
	return &fakeperfdog.UserInfo{
		Name: "user1",
		Id:   "1",
	}, nil
}
