package service

import (
	grpc_logging "github.com/grpc-ecosystem/go-grpc-middleware/logging"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func getUnaryLogger() grpc.UnaryServerInterceptor {
	entry := logrus.NewEntry(logrus.StandardLogger())
	opts := []grpc_logrus.Option{
		grpc_logrus.WithCodes(grpc_logging.DefaultErrorToCode),
		grpc_logrus.WithLevels(grpc_logrus.DefaultClientCodeToLevel),
		grpc_logrus.WithDurationField(grpc_logrus.DefaultDurationToField),
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
	}
	grpc_logrus.ReplaceGrpcLogger(entry)
	return grpc_logrus.StreamServerInterceptor(entry, opts...)
}
