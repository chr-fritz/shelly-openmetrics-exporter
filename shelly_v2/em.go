package shelly_v2

import "github.com/cimnine/shelly-openmetrics-exporter/shelly"

type EMGetConfigRequest struct {
	Id int `json:"id"`
}
type EMGetStatusRequest struct {
	Id int `json:"id"`
}

type EMGetStatusResponse struct {
	// ID of the EM component instance
	Id int `json:"id"`
	// Phase A current measurement value, [A]
	ACurrent float64 `json:"a_current"`
	// Phase A voltage measurement value, [V]
	AVoltage float64 `json:"a_voltage"`
	// Phase A active power measurement value, [W]
	AActivePower float64 `json:"a_act_power"`
	// Phase A apparent power measurement value, [VA]
	AApparentPower float64 `json:"a_aprt_power"`
	// Phase A power factor measurement value
	APowerFactor float64 `json:"a_pf"`
	// Phase A network frequency measurement value
	AFrequency float64 `json:"a_freq"`
	// Phase B current measurement value, [A]
	BCurrent float64 `json:"b_current"`
	// Phase B voltage measurement value, [V]
	BVoltage float64 `json:"b_voltage"`
	// Phase B active power measurement value, [W]
	BActivePower float64 `json:"b_act_power"`
	// Phase B apparent power measurement value, [VA]
	BApparentPower float64 `json:"b_aprt_power"`
	// Phase B power factor measurement value
	BPowerFactor float64 `json:"b_pf"`
	// Phase B network frequency measurement value
	BFrequency float64 `json:"b_freq"`
	// Phase C current measurement value, [A]
	CCurrent float64 `json:"c_current"`
	// Phase C voltage measurement value, [V]
	CVoltage float64 `json:"c_voltage"`
	// Phase C active power measurement value, [W]
	CActivePower float64 `json:"c_act_power"`
	// Phase C apparent power measurement value, [VA]
	CApparentPower float64 `json:"c_aprt_power"`
	// Phase C power factor measurement value
	CPowerFactor float64 `json:"c_pf"`
	// Phase C network frequency measurement value
	CFrequency float64 `json:"c_freq"`
	// Neutral current measurement value, [A] (if supported)
	NCurrent float64 `json:"n_current"`
	// Sum of the current on all phases(excluding neutral readings if available)
	TotalCurrent float64 `json:"total_current"`
	// Sum of the active power on all phases
	TotalActivePower float64 `json:"total_act_power"`
	// Sum of the apparent power on all phases
	TotalApparentPower float64 `json:"total_aprt_power"`
}

type EMGetConfigResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func getEmName(response EMGetConfigResponse) string {
	return response.Name
}

func (s *ShellyV2) fillEMMetrics(m *shelly.Metrics) {
	if s.status.EMStatus == nil {
		return
	}

	for i, measurement := range s.status.EMStatus {
		labelsA := shelly.NamedLineLabels(s.Shelly, "meter", i, shelly.GetConfigValue(s.status.EMConfig, i, getEmName)+": Phase L1")
		m.Voltage.WithLabelValues(labelsA...).Add(measurement.AVoltage)
		m.Current.WithLabelValues(labelsA...).Add(measurement.ACurrent)
		m.Power.WithLabelValues(labelsA...).Add(measurement.AActivePower)
		m.ApparentPower.WithLabelValues(labelsA...).Add(measurement.AApparentPower)
		m.PowerFactor.WithLabelValues(labelsA...).Add(measurement.APowerFactor)
		m.Frequency.WithLabelValues(labelsA...).Add(measurement.AFrequency)

		labelsB := shelly.NamedLineLabels(s.Shelly, "meter", i, shelly.GetConfigValue(s.status.EMConfig, i, getEmName)+": Phase L2")
		m.Voltage.WithLabelValues(labelsB...).Add(measurement.BVoltage)
		m.Current.WithLabelValues(labelsB...).Add(measurement.BCurrent)
		m.Power.WithLabelValues(labelsB...).Add(measurement.BActivePower)
		m.ApparentPower.WithLabelValues(labelsB...).Add(measurement.BApparentPower)
		m.PowerFactor.WithLabelValues(labelsB...).Add(measurement.BPowerFactor)
		m.Frequency.WithLabelValues(labelsB...).Add(measurement.BFrequency)

		labelsC := shelly.NamedLineLabels(s.Shelly, "meter", i, shelly.GetConfigValue(s.status.EMConfig, i, getEmName)+": Phase L3")
		m.Voltage.WithLabelValues(labelsC...).Add(measurement.CVoltage)
		m.Current.WithLabelValues(labelsC...).Add(measurement.CCurrent)
		m.Power.WithLabelValues(labelsC...).Add(measurement.CActivePower)
		m.ApparentPower.WithLabelValues(labelsC...).Add(measurement.CApparentPower)
		m.PowerFactor.WithLabelValues(labelsC...).Add(measurement.CPowerFactor)
		m.Frequency.WithLabelValues(labelsC...).Add(measurement.CFrequency)

		labelsTotal := shelly.NamedLineLabels(s.Shelly, "meter", i, shelly.GetConfigValue(s.status.EMConfig, i, getEmName)+": Total")
		m.Current.WithLabelValues(labelsTotal...).Add(measurement.TotalCurrent)
		m.Power.WithLabelValues(labelsTotal...).Add(measurement.TotalActivePower)
		m.ApparentPower.WithLabelValues(labelsTotal...).Add(measurement.TotalApparentPower)

		labelsN := shelly.NamedLineLabels(s.Shelly, "meter", i, shelly.GetConfigValue(s.status.EMConfig, i, getEmName)+": Neutral")
		m.Current.WithLabelValues(labelsN...).Add(measurement.NCurrent)
	}
}

// getEMStatus retrieves power measurement metrics from the EM component.
// See https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/EM for additional information
func (s *ShellyV2) getEMStatus(status *Status) error {
	for i := 0; true; i++ {
		res := EMGetStatusResponse{}
		request := JsonRpc2Request{
			JsonRpcVersion: "2.0",
			Src:            "shelly-openmetrics-exporter",
			Method:         "EM.GetStatus",
			Params:         EMGetStatusRequest{Id: i},
		}

		end, err := s.do(request, &res)
		if end {
			break
		}
		if err != nil {
			return err
		}

		status.EMStatus = append(status.EMStatus, res)
	}
	return nil
}
func (s *ShellyV2) getEMConfig(status *Status) error {
	for i := 0; true; i++ {
		res := EMGetConfigResponse{}
		request := JsonRpc2Request{
			JsonRpcVersion: "2.0",
			Src:            "shelly-openmetrics-exporter",
			Method:         "EM.GetConfig",
			Params:         EMGetConfigRequest{Id: i},
		}

		end, err := s.do(request, &res)
		if end {
			break
		}
		if err != nil {
			return err
		}

		status.EMConfig = append(status.EMConfig, res)
	}
	return nil
}
