package rectangleArea

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	a1, a2 := (ax2-ax1)*(ay2-ay1), (bx2-bx1)*(by2-by1)
	union := max(min(ax2, bx2)-max(ax1, bx1), 0) * max(min(ay2, by2)-max(ay1, by1), 0)
	//fmt.Println(union)
	return a1 + a2 - union
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
