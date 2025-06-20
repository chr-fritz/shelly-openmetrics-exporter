package shelly_v2

// boolToMetric converts a boolean value either into float64 `1` or `0`
func boolToMetric(value bool) float64 {
	if value {
		return 1
	}
	return 0
}
