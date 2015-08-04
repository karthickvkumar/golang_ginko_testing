package eur

func CurrencyEur(code string, amnt float64) float64 {
	switch code {
	case "INR":
		var ir float64
		ir = amnt * 0.014
		return ir
	case "USD":
		var eu float64
		eu = amnt * 0.90
		return eu
	case "RUB":
		var rb float64
		rb = amnt * 0.016
		return rb
	case "MXN":
		var mx float64
		mx = amnt * 0.057
		return mx
	case "SAR":
		var sa float64
		sa = amnt * 0.24
		return sa
	default:
		return 0.0
	}

}
