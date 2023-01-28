package api

import (
	grpc_logging "github.com/grpc-ecosystem/go-grpc-middleware/logging"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func getUnaryMiddleware() grpc.ServerOption {
	return grpc.ChainUnaryInterceptor(
		getUnaryLogger(),
		grpc_recovery.UnaryServerInterceptor(),
	)
}

func getStreamMiddleware() grpc.ServerOption {
	return grpc.ChainStreamInterceptor(
		getStreamLogger(),
		grpc_recovery.StreamServerInterceptor(),
	)
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
