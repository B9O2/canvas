package containers

import (
	"fmt"
	"testing"
)

var ranges = []Range{
	{
		0,
		0,
	},
	{
		0,
		0,
	},
	{
		0,
		0,
	},
	{
		26,
		26,
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
