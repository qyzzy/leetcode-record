package perfectRectangle

import "math"

func isRectangleCover(rectangles [][]int) bool {
	type ind struct{ x, y int }
	cnt := map[ind]int{}
	area := 0
	minX, minY, maxX, maxY := math.MaxInt32, math.MaxInt32, math.MinInt32, math.MinInt32
	for i := 0; i < len(rectangles); i++ {
		minX = min(minX, rectangles[i][0])
		maxX = max(maxX, rectangles[i][2])
		minY = min(minY, rectangles[i][1])
		maxY = max(maxY, rectangles[i][3])
		area += (rectangles[i][2] - rectangles[i][0]) * (rectangles[i][3] - rectangles[i][1])
		cnt[ind{rectangles[i][0], rectangles[i][1]}]++
		cnt[ind{rectangles[i][0], rectangles[i][3]}]++
		cnt[ind{rectangles[i][2], rectangles[i][1]}]++
		cnt[ind{rectangles[i][2], rectangles[i][3]}]++
	}
	// prevent two rectangles have the same index in border
	if cnt[ind{minX, minY}] != 1 || cnt[ind{maxX, minY}] != 1 || cnt[ind{minX, maxY}] != 1 || cnt[ind{maxX, maxY}] != 1 {
		return false
	}
	if area != (maxX-minX)*(maxY-minY) {
		return false
	}
	// delete the counts of the border's index
	delete(cnt, ind{minX, minY})
	delete(cnt, ind{minX, maxY})
	delete(cnt, ind{maxX, minY})
	delete(cnt, ind{maxX, maxY})
	//fmt.Println(cnt)
	for _, v := range cnt { // two rectangles or four rectangles have the same border's index
		if v > 4 || v%2 == 1 {
			return false
		}
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
