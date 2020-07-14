package challenges

import (
	"fmt"
	"math"
)

// AngleClock -
func AngleClock(hour int, minutes int) float64 {
	angle := math.Abs(float64(float64(30*hour) + (float64(minutes) / 2)) - float64(minutes * 6))
	return math.Min(angle, 360-angle)
}
