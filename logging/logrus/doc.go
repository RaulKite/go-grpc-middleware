// Copyright 2017 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

/*
`grpc_logrus` is a gRPC logging middleware backed by Logrus loggers

It accepts a user-configured `logrus.Entry` that will be used for logging completed gRPC calls. The same
`logrus.Entry` will be used for logging completed gRPC calls, and be populated into the `context.Context` passed into gRPC handler code.

You can use `Extract` to log into a request-scoped `logrus.Entry` instance in your handler code. The fields set on the
logger correspond to the grpc_ctxtags.Tags attached to the context.

Logrus can also be made as a backend for gRPC library internals. For that use `ReplaceGrpcLogger`.

Please see examples and tests for examples of use.
*/
package grpc_logrus
