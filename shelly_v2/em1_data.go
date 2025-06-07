package shelly_v2

import "github.com/cimnine/shelly-openmetrics-exporter/shelly"

type EM1DataGetStatusRequest struct {
	Id int `json:"id"`
}

type EM1DataGetStatusResponse struct {
	Id                   int     `json:"id"`
	ActiveEnergy         float64 `json:"total_act_energy"`
	ActiveReturnedEnergy float64 `json:"total_act_ret_energy"`
}

func (s *ShellyV2) fillEM1DataMetrics(m *shelly.Metrics) {
	if s.status.EM1DataStatus == nil {
		return
	}

	for i, measurement := range s.status.EM1DataStatus {
		labels := shelly.NamedLineLabels(s.Shelly, "meter", i, shelly.GetConfigValue(s.status.EM1Config, i, getEm1Name))

		m.Total.WithLabelValues(labels...).Add(measurement.ActiveEnergy)
		m.TotalReturned.WithLabelValues(labels...).Add(measurement.ActiveReturnedEnergy)
	}
}

// getEM1DataStatus retrieves power measurement metrics from the EM1Data component.
// See https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/EM1Data for additional information
func (s *ShellyV2) getEM1DataStatus(status *Status) error {
	for i := 0; true; i++ {
		res := EM1DataGetStatusResponse{}
		request := JsonRpc2Request{
			JsonRpcVersion: "2.0",
			Src:            "shelly-openmetrics-exporter",
			Method:         "EM1Data.GetStatus",
			Params:         EM1DataGetStatusRequest{Id: i},
		}

		end, err := s.do(request, &res)
		if end {
			break
		}
		if err != nil {
			return err
		}

		status.EM1DataStatus = append(status.EM1DataStatus, res)
	}
	return nil
}
