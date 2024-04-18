package containers

import (
	"fmt"
	"testing"
)

var ranges = []Range{
	{
		36,
		0,
	},
	{
		34,
		0,
	},
	{
		16,
		0,
	},
	{
		0,
		0,
	},
}

func TestSpaceAllocate(t *testing.T) {
	results, err := SpaceAllocate(ranges, 100)
	if err != nil {
		fmt.Println(err)
		return

	}
	for _, r := range results {
		fmt.Println(r)
	}
}
