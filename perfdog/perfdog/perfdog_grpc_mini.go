package perfdog

import (
	context "context"

	grpc "google.golang.org/grpc"
)

type PerfDogServiceServerMini interface {
	LoginWithToken(context.Context, *Token) (*UserInfo, error)
}

// PerfDogServiceMini_ServiceDesc is the grpc.ServiceDesc for PerfDogService service.
var PerfDogServiceMini_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.perfdog.proto.PerfDogService",
	HandlerType: (*PerfDogServiceServerMini)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "loginWithToken",
			Handler:    _PerfDogService_LoginWithToken_Handler_Mini,
		},
	},
	Metadata: "perfdog/perfdog/perfdog.proto",
}

func _PerfDogService_LoginWithToken_Handler_Mini(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerfDogServiceServerMini).LoginWithToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.perfdog.proto.PerfDogService/loginWithToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerfDogServiceServerMini).LoginWithToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func RegisterPerfDogServiceServerMini(s grpc.ServiceRegistrar, srv PerfDogServiceServerMini) {
	s.RegisterService(&PerfDogServiceMini_ServiceDesc, srv)
}
