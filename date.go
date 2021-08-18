package aphgrpc

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ProtoTimeStamp(ts *timestamppb.Timestamp) time.Time {
	return ts.AsTime()
}

func TimestampProto(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}
