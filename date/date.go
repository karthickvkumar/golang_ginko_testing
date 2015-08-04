package date

import "time"

func Years() int {
	now := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	return now.Year()
}
func Format() string {
	t := time.Now()
	return t.Format("3 PM")
}
