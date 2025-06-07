package shelly_v2

import (
	"github.com/cimnine/shelly-openmetrics-exporter/shelly"
)

type Status struct {
	SwitchesStatus    []SwitchGetStatusResponse
	SwitchesConfig    []SwitchGetConfigResponse
	InputStatus       []InputGetStatusResponse
	WifiStatus        *WifiGetStatusResponse
	CloudStatus       *CloudGetStatusResponse
	CloudConfig       *CloudGetConfigResponse
	VoltmeterStatus   []VoltmeterGetStatusResponse
	TemperatureStatus []TemperatureGetStatusResponse
	HumidityStatus    []HumidityGetStatusResponse
	DevicePowerStatus []DevicePowerGetStatusResponse
	LightStatus       []LightGetStatusResponse
	CoversConfig      []CoverGetConfigResponse
	CoversStatus      []CoverGetStatusResponse
	PM1Status         []PM1GetStatusResponse
	PM1Config         []PM1GetConfigResponse
	EM1Status         []EM1GetStatusResponse
	EM1Config         []EM1GetConfigResponse
	EM1DataStatus     []EM1DataGetStatusResponse
}

func (s *ShellyV2) FillMetrics(m *shelly.Metrics) {
	s.fillSwitchMetrics(m)
	s.fillInputMetrics(m)
	s.fillWifiMetrics(m)
	s.fillCloudMetrics(m)
	s.fillVoltmeterMetrics(m)
	s.fillHumidityMetrics(m)
	s.fillTemperatureMetrics(m)
	s.fillDevicePowerMetrics(m)
	s.fillCoverMetrics(m)
	s.fillPM1Metrics(m)
	s.fillEM1Metrics(m)
	s.fillEM1DataMetrics(m)
}
