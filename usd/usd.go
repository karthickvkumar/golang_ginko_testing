package usd

func CurrencyUsd(code string, amnt float64) float64 {
	switch code {
	case "INR":
		var ir float64
		ir = amnt * 0.016
		return ir
	case "EUR":
		var eu float64
		eu = amnt * 1.11
		return eu
	case "RUB":
		var rb float64
		rb = amnt * 62.56
		return rb
	case "MXN":
		var mx float64
		mx = amnt * 17.46
		return mx
	case "SAR":
		var sa float64
		sa = amnt * 4.17
		return sa
	default:
		return 0.0
	}

}
