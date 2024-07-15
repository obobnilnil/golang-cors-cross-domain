package utility

import "fmt"

func FormatK(value float64) string {
	if value >= 1000 {
		return fmt.Sprintf("%.1fK", value/1000)
	}
	return fmt.Sprintf("%.0f", value)
}
