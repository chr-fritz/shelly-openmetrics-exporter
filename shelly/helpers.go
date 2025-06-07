package shelly

import (
	"strconv"
)

func BoolToFloat(b bool) float64 {
	if b {
		return 1
	}

	return 0
}

func DeviceLabels(s *Shelly) []string {
	return []string{s.TargetHost}
}

func LineLabels(s *Shelly, kind string, line int) []string {
	return append(DeviceLabels(s), kind+":"+strconv.Itoa(line))
}

func NamedLineLabels(s *Shelly, kind string, line int, name string) []string {
	return append(DeviceLabels(s), kind+":"+strconv.Itoa(line), name)
}

func CelsiusToKelvin(celsius float64) float64 {
	return celsius + zeroCelsiusInKelvin
}

const zeroCelsiusInKelvin = 273.15

func GetConfigValue[T any](configs []T, index int, getter func(T) string) string {
	if len(configs) == 0 || len(configs) <= index {
		return ""
	}
	return getter(configs[index])
}
