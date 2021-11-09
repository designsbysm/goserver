package collatz

import "github.com/designsbysm/server-go/rpc/collatzpb"

func Hailstone(seed int64) (*collatzpb.SeedResponse, error) {
	var result collatzpb.SeedResponse
	var err error

	result.Value = seed
	result.Path = append(result.Path, seed)

	for seed > 1 {
		seed, err = stone(seed)
		if err != nil {
			return &result, err
		}

		result.Path = append(result.Path, seed)
	}

	return &result, nil
}
