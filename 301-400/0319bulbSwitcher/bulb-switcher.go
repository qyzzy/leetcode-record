package bulbSwitcher

import "math"

func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n) + 0.5))
}
