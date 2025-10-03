package shelly_v2

import "github.com/cimnine/shelly-openmetrics-exporter/shelly"

type EMDataGetStatusRequest struct {
	Id int `json:"id"`
}

type EMDataGetStatusResponse struct {
	Id int `json:"id"`
	// Total active energy on phase A, Wh
	AActiveEnergy float64 `json:"a_total_act_energy"`
	// Total active returned energy on phase A, Wh
	AActiveReturnedEnergy float64 `json:"a_total_act_ret_energy"`
	// Total active energy on phase B, Wh
	BActiveEnergy float64 `json:"b_total_act_energy"`
	// Total active returned energy on phase B, Wh
	BActiveReturnedEnergy float64 `json:"b_total_act_ret_energy"`
	// Total active energy on phase C, Wh
	CActiveEnergy float64 `json:"c_total_act_energy"`
	// Total active returned energy on phase C, Wh
	CActiveReturnedEnergy float64 `json:"c_total_act_ret_energy"`
	// Total active energy on all phases, Wh
	TotalActiveEnergy float64 `json:"total_act"`
	// Total active returned energy on all phases, Wh
	TotalActiveReturnedEnergy float64 `json:"total_act_ret"`
}

func (s *ShellyV2) fillEMDataMetrics(m *shelly.Metrics) {
	if s.status.EMDataStatus == nil {
		return
	}

	for i, measurement := range s.status.EMDataStatus {
		labelsA := shelly.NamedLineLabels(s.Shelly, "meter", i, shelly.GetConfigValue(s.status.EMConfig, i, getEmName)+": Phase L1")
		m.Total.WithLabelValues(labelsA...).Add(measurement.AActiveEnergy)
		m.TotalReturned.WithLabelValues(labelsA...).Add(measurement.AActiveReturnedEnergy)

		labelsB := shelly.NamedLineLabels(s.Shelly, "meter", i, shelly.GetConfigValue(s.status.EMConfig, i, getEmName)+": Phase L2")
		m.Total.WithLabelValues(labelsB...).Add(measurement.BActiveEnergy)
		m.TotalReturned.WithLabelValues(labelsB...).Add(measurement.BActiveReturnedEnergy)

		labelsC := shelly.NamedLineLabels(s.Shelly, "meter", i, shelly.GetConfigValue(s.status.EMConfig, i, getEmName)+": Phase L3")
		m.Total.WithLabelValues(labelsC...).Add(measurement.CActiveEnergy)
		m.TotalReturned.WithLabelValues(labelsC...).Add(measurement.CActiveReturnedEnergy)

		labelsTotal := shelly.NamedLineLabels(s.Shelly, "meter", i, shelly.GetConfigValue(s.status.EMConfig, i, getEmName)+": Total")
		m.Total.WithLabelValues(labelsTotal...).Add(measurement.TotalActiveEnergy)
		m.TotalReturned.WithLabelValues(labelsTotal...).Add(measurement.TotalActiveReturnedEnergy)
	}
}

// getEMDataStatus retrieves power measurement metrics from the EMData component.
// See https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/EMData for additional information
func (s *ShellyV2) getEMDataStatus(status *Status) error {
	for i := 0; true; i++ {
		res := EMDataGetStatusResponse{}
		request := JsonRpc2Request{
			JsonRpcVersion: "2.0",
			Src:            "shelly-openmetrics-exporter",
			Method:         "EMData.GetStatus",
			Params:         EMDataGetStatusRequest{Id: i},
		}

		end, err := s.do(request, &res)
		if end {
			break
		}
		if err != nil {
			return err
		}

		status.EMDataStatus = append(status.EMDataStatus, res)
	}
	return nil
}
