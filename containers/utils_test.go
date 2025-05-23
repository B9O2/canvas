package containers

import (
	"fmt"
	"reflect"
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

func TestSplitWithLengthUnicode(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		length   uint
		expected []string
	}{
		{"Chinese basic", "你好世界", 4, []string{"你好", "世界"}},
		{"Chinese uneven split 1", "你好世界", 3, []string{"你", "好", "世", "界"}}, // Each char is width 2. "你"(2) fits. "好"(2) overflows 3-2=1.
		{"Chinese uneven split 2", "你好世界", 2, []string{"你", "好", "世", "界"}},
		{"ASCII basic", "abcde", 2, []string{"ab", "cd", "e"}},
		{"Emoji wider than length", "🚀", 1, []string{"🚀"}}, // Rocket emoji width 2
		{"Emoji fits length", "🚀", 2, []string{"🚀"}},
		{"Emoji then ASCII", "🚀abc", 3, []string{"🚀a", "bc"}},    // Corrected: 🚀 (2)+a(1)=3. Next 'b' overflows.
		{"Emoji then ASCII overflow", "🚀abc", 2, []string{"🚀", "ab", "c"}}, // 🚀 (2). a(1)+b(1)=2. c(1) overflows. This one was correct.
		{"ASCII then Emoji", "abc🚀", 3, []string{"abc", "🚀"}},    // abc (3). 🚀 (2) overflows.
		{"ASCII then Emoji fit", "abc🚀", 5, []string{"abc🚀"}}, // abc (3) + 🚀 (2) = 5
		{"ASCII then Emoji partial fit", "abc🚀", 4, []string{"abc", "🚀"}}, // abc (3). Remaining 1. 🚀 (2) overflows.
		{"Empty string", "", 5, []string{}}, // Behavior for empty string with length > 0
		{"Non-empty string zero length", "test", 0, []string{"test"}}, // Current behavior for 0 length
		{"Chinese string zero length", "你好", 0, []string{"你好"}},     // Current behavior for 0 length
		{"Mixed string", "a你好b世界c", 6, []string{"a你好b", "世界c"}}, // Corrected: a(1)你好(4)b(1)=6. Next '世' overflows.
		{"Mixed string 2", "a你好b世界c", 7, []string{"a你好b", "世界c"}}, // This one was correct: a(1)+你(2)+好(2)+b(1)=6. Next '世'(2) overflows. Remaining '世界c'.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SplitWithLength(tt.text, tt.length)
			// Using reflect.DeepEqual for slice comparison
			// For empty slices, DeepEqual handles nil vs non-nil empty slices correctly.
			// However, SplitWithLength initializes result with `var result []string` which is nil.
			// If expected is `[]string{}`, it's a non-nil empty slice.
			// Let's ensure expected empty is also nil for consistency if result is nil.
			if len(tt.expected) == 0 && len(result) == 0 {
				// Both are effectively empty, pass without DeepEqual if one is nil and other is empty non-nil
			} else if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("SplitWithLength(%q, %d) = %v, want %v", tt.text, tt.length, result, tt.expected)
			}
		})
	}
}
