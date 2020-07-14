package challenges

import (
	"fmt"
	"math"
)

// AngleClock -
func AngleClock(hour int, minutes int) float64 {
	angle1 := math.Abs(getHourAngle(hour, minutes) - getMinuteAngle(minutes))
	angle2 := angle1
	angle2 = 360 - angle2
	return math.Min(angle1, angle2)
}
func getHourAngle(hour int, minutes int) float64 {
	fmt.Printf("Hour angle is %f", float64(float64(30*hour)+(float64(minutes)/2)))
	angle := float64(float64(30*hour) + (float64(minutes) / 2))
	if angle > 360 {
		angle -= 360
	}
	return angle

}
func getMinuteAngle(minutes int) float64 {
	fmt.Printf("Minutes angle is %f", float64(minutes*6))

	return float64(minutes * 6)
}
