package round

import "math"

func Rounding(code string, value float64) float64 {
	switch code {
	case "Floor":
		var f float64
		f = math.Floor(value)
		return f
	case "Ceil":
		var c float64
		c = math.Ceil(value)
		return c
	default:
		return 0.0
	}

}
