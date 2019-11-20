// Package aphgrpc provides various interfaces, functions, types
// for building and working with gRPC services.
package aphgrpc

import (
	"context"
	"errors"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	// MetaKey is the key used for storing all metadata
	MetaKey = "error"
)

var (
	//ErrDatabaseQuery represents database query related errors
	ErrDatabaseQuery = newError("Database query error")
	//ErrDatabaseInsert represents database insert related errors
	ErrDatabaseInsert = newError("Database insert error")
	//ErrDatabaseUpdate represents database update related errors
	ErrDatabaseUpdate = newError("Database update error")
	//ErrDatabaseDelete represents database update delete errors
	ErrDatabaseDelete = newError("Database delete error")
	//ErrNotFound represents the absence of an HTTP resource
	ErrNotFound = newError("Resource not found")
	//ErrExists represents the presence of an HTTP resource
	ErrExists = newError("Resource already exists")
	//ErrJSONEncoding represents any json encoding error
	ErrJSONEncoding = newError("Json encoding error")
	//ErrStructMarshal represents any error with marshalling structure
	ErrStructMarshal = newError("Structure marshalling error")
	//ErrIncludeParam represents any error with invalid include query parameter
	ErrIncludeParam = newErrorWithParam("Invalid include query parameter", "include")
	//ErrSparseFieldSets represents any error with invalid sparse fieldsets query parameter
	ErrFields = newErrorWithParam("Invalid field query parameter", "field")
	//ErrFilterParam represents any error with invalid filter query paramter
	ErrFilterParam = newErrorWithParam("Invalid filter query parameter", "filter")
	//ErrNotAcceptable represents any error with wrong or inappropriate http Accept header
	ErrNotAcceptable = newError("Accept header is not acceptable")
	//ErrUnsupportedMedia represents any error with unsupported media type in http header
	ErrUnsupportedMedia = newError("Media type is not supported")
	//ErrInValidParam represents any error with validating input parameters
	ErrInValidParam = newError("Invalid parameters")
	//ErrRetrieveMetadata represents any error to retrieve grpc metadata from the running context
	ErrRetrieveMetadata = errors.New("unable to retrieve metadata")
	//ErrXForwardedHost represents any failure or absence of x-forwarded-host HTTP header in the grpc context
	ErrXForwardedHost = errors.New("x-forwarded-host header is absent")
)

func newErrorWithParam(msg, param string) metadata.MD {
	return metadata.Pairs(MetaKey, msg, MetaKey, param)
}

func newError(msg string) metadata.MD {
	return metadata.Pairs(MetaKey, msg)
}

func getgRPCStatus(err error) *status.Status {
	s, ok := status.FromError(err)
	if !ok {
		return status.New(codes.Unknown, err.Error())
	}
	return s
}

func CheckNoRows(err error) bool {
	if strings.Contains(err.Error(), "no rows") {
		return true
	}
	return false
}

func HandleError(ctx context.Context, err error) error {
	if CheckNoRows(err) {
		grpc.SetTrailer(ctx, ErrNotFound)
		return status.Error(codes.NotFound, err.Error())
	}
	grpc.SetTrailer(ctx, newError(err.Error()))
	return status.Error(codes.Internal, err.Error())
}

func HandleGenericError(ctx context.Context, err error) error {
	grpc.SetTrailer(ctx, newError(err.Error()))
	return status.Error(codes.Internal, err.Error())
}

func HandleDeleteError(ctx context.Context, err error) error {
	grpc.SetTrailer(ctx, ErrDatabaseDelete)
	return status.Error(codes.Internal, err.Error())
}

func HandleGetError(ctx context.Context, err error) error {
	grpc.SetTrailer(ctx, ErrDatabaseQuery)
	return status.Error(codes.Internal, err.Error())
}

func HandleInsertError(ctx context.Context, err error) error {
	grpc.SetTrailer(ctx, ErrDatabaseInsert)
	return status.Error(codes.Internal, err.Error())
}

func HandleUpdateError(ctx context.Context, err error) error {
	grpc.SetTrailer(ctx, ErrDatabaseUpdate)
	return status.Error(codes.Internal, err.Error())
}

func HandleGetArgError(ctx context.Context, err error) error {
	grpc.SetTrailer(ctx, ErrDatabaseQuery)
	return status.Error(codes.InvalidArgument, err.Error())
}

func HandleInsertArgError(ctx context.Context, err error) error {
	grpc.SetTrailer(ctx, ErrDatabaseInsert)
	return status.Error(codes.InvalidArgument, err.Error())
}

func HandleUpdateArgError(ctx context.Context, err error) error {
	grpc.SetTrailer(ctx, ErrDatabaseUpdate)
	return status.Error(codes.InvalidArgument, err.Error())
}

func HandleNotFoundError(ctx context.Context, err error) error {
	grpc.SetTrailer(ctx, ErrNotFound)
	return status.Error(codes.NotFound, err.Error())
}

func HandleExistError(ctx context.Context, err error) error {
	grpc.SetTrailer(ctx, ErrExists)
	return status.Error(codes.AlreadyExists, err.Error())
}

func HandleFilterParamError(ctx context.Context, err error) error {
	grpc.SetTrailer(ctx, ErrFilterParam)
	return status.Error(codes.InvalidArgument, err.Error())
}

func HandleInvalidParamError(ctx context.Context, err error) error {
	grpc.SetTrailer(ctx, ErrInValidParam)
	return status.Error(codes.InvalidArgument, err.Error())
}
