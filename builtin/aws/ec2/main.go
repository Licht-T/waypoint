package ec2

import (
	"github.com/hashicorp/waypoint/sdk"
)

//go:generate protoc -I ../../../.. --go_opt=plugins=grpc --go_out=../../../.. waypoint/builtin/aws/ec2/plugin.proto

// Options are the SDK options to use for instantiation.
var Options = []sdk.Option{
	sdk.WithComponents(&Platform{}),
}