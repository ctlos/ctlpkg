package api

import (
	"context"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_logging "github.com/grpc-ecosystem/go-grpc-middleware/logging"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	authType   = "Bearer"
	dummyToken = "12345"
)

var ErrAuthFailed = status.Errorf(codes.PermissionDenied, "buildDummyAuthFunction bad token")

func getUnaryMiddleware() grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(
		getUnaryLogger(),
		grpc_auth.UnaryServerInterceptor(auth),
		grpc_recovery.UnaryServerInterceptor(),
	)
}

func getStreamMiddleware() grpc.ServerOption {
	return grpc.ChainStreamInterceptor(
		getStreamLogger(),
		grpc_auth.StreamServerInterceptor(auth),
		grpc_recovery.StreamServerInterceptor(),
	)
}

func auth(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, authType)
	if err != nil {
		return nil, err
	}
	if token != dummyToken {
		return nil, ErrAuthFailed
	}
	return ctx, nil
}

func getUnaryLogger() grpc.UnaryServerInterceptor {
	entry := logrus.NewEntry(logrus.StandardLogger())
	opts := []grpc_logrus.Option{
		grpc_logrus.WithCodes(grpc_logging.DefaultErrorToCode),
		grpc_logrus.WithLevels(grpc_logrus.DefaultClientCodeToLevel),
		grpc_logrus.WithDurationField(grpc_logrus.DefaultDurationToField),
		grpc_logrus.WithMessageProducer(grpc_logrus.DefaultMessageProducer),
	}
	grpc_logrus.ReplaceGrpcLogger(entry)
	return grpc_logrus.UnaryServerInterceptor(entry, opts...)
}

func getStreamLogger() grpc.StreamServerInterceptor {
	entry := logrus.NewEntry(logrus.StandardLogger())
	opts := []grpc_logrus.Option{
		grpc_logrus.WithCodes(grpc_logging.DefaultErrorToCode),
		grpc_logrus.WithLevels(grpc_logrus.DefaultClientCodeToLevel),
		grpc_logrus.WithDurationField(grpc_logrus.DefaultDurationToField),
		grpc_logrus.WithMessageProducer(grpc_logrus.DefaultMessageProducer),
	}
	grpc_logrus.ReplaceGrpcLogger(entry)
	return grpc_logrus.StreamServerInterceptor(entry, opts...)
}
