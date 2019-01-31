package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	maxSize, safetySize := compute(data, 10000)
	fmt.Printf("maxSize: %d, safetySize:%d\n", maxSize, safetySize)
}

func compute(data string, safety_distance int) (int, int) {
	scanner := bufio.NewScanner(strings.NewReader(data))
	points := make([][]int, 0)
	minx := math.MaxInt64
	miny := math.MaxInt64
	maxx := 0
	maxy := 0

	for scanner.Scan() {
		line := scanner.Text()
		coord := strings.Split(line, ",")
		x, err := strconv.Atoi(strings.TrimSpace(coord[0]))
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(strings.TrimSpace(coord[1]))
		if err != nil {
			panic(err)
		}

		points = append(points, []int{x, y})
		minx = min(x, minx)
		miny = min(y, miny)
		maxx = max(x, maxx)
		maxy = max(y, maxy)
	}
	areas := make([]int, len(points))
	// set area extrems
	minx = minx - 1
	miny = miny - 1
	maxx = maxx + 1
	maxy = maxy + 0

	fmt.Printf("grid size is (%d,%d) (%d,%d)\n", minx, miny, maxx, maxy)

	safety_size := 0
	exclusions := make([]int, 0)
	for i := minx; i <= maxx; i++ {
		for j := miny; j <= maxy; j++ {
			curr_point := []int{i, j}
			mindist := math.MaxInt64
			minpoints := make([]int, 0)
			manhattan_total := 0
			// looking for min manhattan distance
			for p, c := range points {
				d := manhattan(curr_point, c)
				manhattan_total = manhattan_total + d
				if d < mindist {
					mindist = d
					minpoints = []int{p}
				} else if d == mindist {
					minpoints = append(minpoints, p)
				}
			}
			if manhattan_total < safety_distance {
				safety_size++
			}
			if len(minpoints) == 1 {
				if memberof(i, []int{minx, maxx}) || memberof(j, []int{miny, maxy}) {
					// point is on an edge, so bound area should be excluded
					if !memberof(minpoints[0], exclusions) {
						exclusions = append(exclusions, minpoints[0])
					}
				}
				// point can be added to an area for increasing the area
				areas[minpoints[0]] = areas[minpoints[0]] + 1
			}
		}
	}

	maxArea := 0
	for p, a := range areas {
		if memberof(p, exclusions) {
			continue
		}
		maxArea = max(a, maxArea)
	}
	return maxArea, safety_size
}

func manhattan(p1, p2 []int) int {
	//fmt.Printf("manhattan:(%v;%v) - %d\n", p1, p2, Abs(p1[0]-p2[0])+Abs(p1[1]-p2[1]))
	return Abs(p1[0]-p2[0]) + Abs(p1[1]-p2[1])
}

func Abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func memberof(m int, list []int) bool {
	for _, l := range list {
		if l == m {
			return true
		}
	}
	return false
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

var data = `227, 133
140, 168
99, 112
318, 95
219, 266
134, 144
306, 301
189, 188
58, 334
337, 117
255, 73
245, 144
102, 257
255, 353
303, 216
141, 167
40, 321
201, 50
60, 188
132, 74
125, 199
176, 307
204, 218
338, 323
276, 278
292, 229
109, 228
85, 305
86, 343
97, 254
182, 151
110, 292
285, 124
43, 223
153, 188
285, 136
334, 203
84, 243
92, 185
330, 223
259, 275
106, 199
183, 205
188, 212
231, 150
158, 95
174, 212
279, 97
172, 131
247, 320`
