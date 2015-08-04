package tax

func People(age int) string {

	if age > 0 && age <= 50 {
		return "Normal Citizen"
	} else if age > 50 && age <= 61 {
		return "Seniour Citizen"
	} else {
		return "Super Seniour Citizen"
	}
}
func Taxs(category string, income float64) float64 {
	switch category {
	case "GEN":
		if income > 0 && income <= 250000 {
			return 0.0
		} else if income > 250001 && income <= 500000 {
			return (income - 250000) * (0.1)

		} else if income > 500001 && income <= 1000000 {
			return ((income-250000)*(0.1) + (income-500000)*(0.2))
		} else {
			return ((income-250000)*(0.1) + (income-500000)*(0.2) + (income-1000000)*(0.3))
		}
	case "SEN":
		if income > 0 && income <= 300000 {
			return 0.0
		} else if income > 300001 && income <= 500000 {
			return (income - 300000) * (0.1)

		} else if income > 500001 && income <= 1000000 {
			return ((income-300000)*(0.1) + (income-500000)*(0.2))

		} else {
			return ((income-300000)*(0.1) + (income-500000)*(0.2) + (income-1000000)*(0.3))
		}
	case "SUP_SEN":
		if income > 0 && income <= 500000 {
			return 0.0
		} else if income > 500001 && income <= 1000000 {
			return (income - 500000) * (0.2)
		} else {
			return ((income-500000)*(0.2) + (income-1000000)*(0.3))
		}

	default:
		return 0.0
	}
}
