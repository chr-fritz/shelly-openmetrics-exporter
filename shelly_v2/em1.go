package shelly_v2

import "github.com/cimnine/shelly-openmetrics-exporter/shelly"

type EM1GetConfigRequest struct {
	Id int `json:"id"`
}
type EM1GetStatusRequest struct {
	Id int `json:"id"`
}

type EM1GetStatusResponse struct {
	Id            int     `json:"id"`
	Voltage       float64 `json:"voltage"`
	Current       float64 `json:"current"`
	ActivePower   float64 `json:"act_power"`
	ApparentPower float64 `json:"aprt_power"`
	PowerFactor   float64 `json:"pf"`
}

type EM1GetConfigResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Reverse bool   `json:"reverse"`
	CtType  string `json:"ct_type"`
}

func getEm1Name(response EM1GetConfigResponse) string {
	return response.Name
}

func (s *ShellyV2) fillEM1Metrics(m *shelly.Metrics) {
	if s.status.EM1Status == nil {
		return
	}

	for i, measurement := range s.status.EM1Status {
		labels := shelly.NamedLineLabels(s.Shelly, "meter", i, shelly.GetConfigValue(s.status.EM1Config, i, getEm1Name))

		m.Voltage.WithLabelValues(labels...).Add(measurement.Voltage)
		m.Current.WithLabelValues(labels...).Add(measurement.Current)
		m.Power.WithLabelValues(labels...).Add(measurement.ActivePower)
		m.ApparentPower.WithLabelValues(labels...).Add(measurement.ApparentPower)
		m.PowerFactor.WithLabelValues(labels...).Add(measurement.PowerFactor)
	}
}

// getEM1Status retrieves power measurement metrics from the EM1 component.
// See https://shelly-api-docs.shelly.cloud/gen2/ComponentsAndServices/EM1 for additional information
func (s *ShellyV2) getEM1Status(status *Status) error {
	for i := 0; true; i++ {
		res := EM1GetStatusResponse{}
		request := JsonRpc2Request{
			JsonRpcVersion: "2.0",
			Src:            "shelly-openmetrics-exporter",
			Method:         "EM1.GetStatus",
			Params:         EM1GetStatusRequest{Id: i},
		}

		end, err := s.do(request, &res)
		if end {
			break
		}
		if err != nil {
			return err
		}

		status.EM1Status = append(status.EM1Status, res)
	}
	return nil
}
func (s *ShellyV2) getEM1Config(status *Status) error {
	for i := 0; true; i++ {
		res := EM1GetConfigResponse{}
		request := JsonRpc2Request{
			JsonRpcVersion: "2.0",
			Src:            "shelly-openmetrics-exporter",
			Method:         "EM1.GetConfig",
			Params:         EM1GetConfigRequest{Id: i},
		}

		end, err := s.do(request, &res)
		if end {
			break
		}
		if err != nil {
			return err
		}

		status.EM1Config = append(status.EM1Config, res)
	}
	return nil
}
