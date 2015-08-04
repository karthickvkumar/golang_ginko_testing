package india

func CurrencyInr(code string, amnt float64) float64 {
	switch code {
	case "USD":
		var us float64
		us = amnt * 64.10
		return us
	case "EUR":
		var eu float64
		eu = amnt * 71.14
		return eu
	case "RUB":
		var rb float64
		rb = amnt * 1.13
		return rb
	case "MXN":
		var mx float64
		mx = amnt * 4.08
		return mx
	case "SAR":
		var sa float64
		sa = amnt * 18.0
		return sa
	default:
		return 0.0
	}

}
