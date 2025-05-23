package containers

import (
	"fmt"
	"sort"

	"github.com/B9O2/canvas/pixel"
	"github.com/mattn/go-runewidth"
)

func SpaceAllocate(ranges []Range, length uint) ([]uint, error) {
	type RangeWithRawID struct {
		r        Range
		id       int
		allocate float32
	}

	var idRanges []RangeWithRawID
	for id, r := range ranges {
		idRanges = append(idRanges, RangeWithRawID{
			id: id,
			r:  r,
		})
	}

	total := len(idRanges)
	if total < 1 {
		return []uint{}, nil
	}

	totalMinLength := uint(0)
	totalMaxLength := uint(0)
	hasZeroMax := false
	for i := range idRanges {
		min := idRanges[i].r.Min()
		max := idRanges[i].r.Max()
		totalMinLength += min
		if max == 0 {
			hasZeroMax = true
		}
		totalMaxLength += max
		idRanges[i].allocate = float32(min)
	}
	if totalMinLength > length {
		return nil, fmt.Errorf("space too small")
	}
	if !hasZeroMax && totalMaxLength < length {
		return nil, fmt.Errorf("space too large")
	}

	sort.Slice(idRanges, func(i, j int) bool {
		return idRanges[i].r.Min() < idRanges[j].r.Min()
	})

	//Calculate
	remainingSpace := float32(length - totalMinLength)
	for remainingSpace > 0 {
		for i := range idRanges {
			r := idRanges[i]
			ids := FindNums(idRanges, func(v RangeWithRawID) bool {
				if v.r.Max() > 0 {
					return (v.allocate < r.allocate && v.allocate < float32(v.r.Max()))
				}
				return v.allocate < r.allocate
			})
			//fmt.Println("目标", idRanges[i].allocate, "不符合", ids)
			totalNeedSpace := float32(0)
			need := float32(len(ids))
			if need <= 0 {
				if i >= total-1 { //final
					totalNeedSpace = remainingSpace
					need = float32(total)
					ids = GenNumbers[int](0, total-1)
				} else {
					continue
				}
			} else {
				for _, id := range ids {
					if idRanges[id].r.Max() > 0 && r.allocate > float32(idRanges[id].r.Max()) {
						totalNeedSpace += float32(idRanges[id].r.Max()) - idRanges[id].allocate
					} else {
						totalNeedSpace += r.allocate - idRanges[id].allocate
					}
				}
			}
			//fmt.Println("Need", need, "Total", totalNeedSpace)
			if totalNeedSpace >= remainingSpace {
				space := remainingSpace / need
				//fmt.Println("分配", space)
				remainingSpace = 0
				for _, id := range ids {
					//fmt.Printf("%f -- +%f --> %f \n", idRanges[id].allocate, space, idRanges[id].allocate+space)
					max := float32(idRanges[id].r.Max())
					if max > 0 && idRanges[id].allocate+space > max {
						remainingSpace += idRanges[id].allocate + space - max
						idRanges[id].allocate = max
					} else {
						idRanges[id].allocate += space
					}
				}
			} else {
				for _, id := range ids {
					max := float32(idRanges[id].r.Max())
					less := r.allocate - idRanges[id].allocate
					remainingSpace -= less
					//fmt.Printf("%f -- +%f --> %f Remaining: %f\n", idRanges[id].allocate, less, idRanges[id].allocate+less, remainingSpace)
					if max > 0 && r.allocate > max {
						remainingSpace += r.allocate - max
						idRanges[id].allocate = max
					} else {
						idRanges[id].allocate += less
					}
				}
			}
		}
	}

	// sort.Slice(idRanges, func(i, j int) bool {
	// 	return idRanges[i].r.Min() < idRanges[j].r.Min()
	// })
	// remainingSpace := float32(length - totalMinLength)
	// for remainingSpace > 0 {
	// 	fmt.Println("-------------------------")
	// 	for i := float32(1); i <= float32(total); i++ {
	// 		space := remainingSpace / i
	// 		ids := FindNums(idRanges, func(v RangeWithRawID) bool {
	// 			return v.allocate < space
	// 		})
	// 		need := len(ids)
	// 		fmt.Printf("All %v Target %f Need %v \n", idRanges, space, ids)
	// 		if need == 0 {
	// 			fmt.Printf("All Qual %v\n", idRanges)
	// 			space = remainingSpace / float32(total)
	// 			for j := range idRanges {
	// 				fmt.Printf("%f -- +%f --> %f\n", idRanges[j].allocate, space, idRanges[j].allocate+space)
	// 				idRanges[j].allocate += space
	// 				remainingSpace = 0
	// 			}
	// 			break
	// 		} else if need > int(i) { //有不止i个需被分配
	// 			fmt.Printf("Try allocate %f , but %d(%v)\n", i, need, ids)
	// 			continue
	// 		} else {
	// 			for _, id := range ids {
	// 				less := space - idRanges[id].allocate
	// 				remainingSpace -= less
	// 				fmt.Printf("%f -- +%f --> %f Remaining: %f\n", idRanges[id].allocate, less, idRanges[id].allocate+less, remainingSpace)
	// 				idRanges[id].allocate += less
	// 			}
	// 		}
	// 	}
	// }

	/*
		remainingSpace := length - totalMinLength

		for remainingSpace > 0 {
			space := remainingSpace / total
			remainingSpace = 0
			for i := uint(0); i < total; i++ {
				r := ranges[i]
				results[i] += space
				maxLength := r.Max()
				if maxLength > 0 && results[i] > maxLength {
					remainingSpace += (results[i] - maxLength)
					results[i] = maxLength
				}
			}
		}
	*/

	results := make([]uint, total)
	for _, r := range idRanges {
		results[r.id] = uint(r.allocate)
	}

	return results, nil
}

func FindNums[T any](nums []T, f func(v T) bool) []int {
	var results []int
	for id, num := range nums {
		if f(num) {
			results = append(results, id)
		}
	}
	return results
}

func PixelString(text string, color any) []pixel.Pixel {
	str := []rune(text)
	pixels := make([]pixel.Pixel, len(str))
	for i, r := range str {
		pixels[i] = pixel.NewPixel(r, color)
	}
	return pixels
}

func SplitWithLength(text string, length uint) []string {
	var result []string
	if length == 0 {
		if text == "" {
			return []string{}
		}
		return []string{text} // Or handle as an error, or return a list of single-character strings
	}

	runes := []rune(text)
	currentLineWidth := uint(0)
	start := 0
	for i, r := range runes {
		runeWidth := uint(runewidth.RuneWidth(r))

		if runeWidth > length { // Single character wider than limit
			if start < i { // Add preceding characters as a line
				result = append(result, string(runes[start:i]))
			}
			result = append(result, string(r)) // Add wide character as its own line
			currentLineWidth = 0
			start = i + 1
			continue
		}

		if currentLineWidth+runeWidth > length {
			result = append(result, string(runes[start:i]))
			start = i
			currentLineWidth = runeWidth
		} else {
			currentLineWidth += runeWidth
		}
	}

	if start < len(runes) {
		result = append(result, string(runes[start:]))
	}
	return result
}

func GenNumbers[T uint | int | float32 | float64](min, max T) []T {
	a := make([]T, int(max-min+1))
	for i := range a {
		a[i] = min + T(i)
	}
	return a
}
