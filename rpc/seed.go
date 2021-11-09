package rpc

import (
	"context"

	"github.com/designsbysm/server-go/rpc/collatz"
	"github.com/designsbysm/server-go/rpc/collatzpb"
	"github.com/designsbysm/timber/v2"
)

func (*server) Seed(ctx context.Context, in *collatzpb.SeedRequest) (*collatzpb.SeedResponse, error) {
	seed := in.GetValue()
	// return collatz.Hailstone(seed)

	res, err := collatz.Hailstone(seed)

	timber.Struct(res)

	return res, err
}
