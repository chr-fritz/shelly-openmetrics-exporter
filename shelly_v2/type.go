package shelly_v2

import (
	"net/http"

	"github.com/cimnine/shelly-openmetrics-exporter/shelly"
)

type ShellyV2StatusFetcher func(*Status) error
type ShellyV2 struct {
	*shelly.Shelly
	status        *Status
	nextMessageID *int
	client        *http.Client
	fetchers      []ShellyV2StatusFetcher
}

func New(targetHost, userAgent, password string) *ShellyV2 {
	initialMessageId := 0
	api := &ShellyV2{
		Shelly: &shelly.Shelly{
			TargetHost: targetHost,
			UserAgent:  userAgent,
			Username:   "admin", // must be admin, per documentation
			Password:   password,
		},
		nextMessageID: &initialMessageId,
		client:        &http.Client{},
	}
	api.fetchers = []ShellyV2StatusFetcher{
		api.getSwitchStatus,
		api.getSwitchConfig,
		api.getInputStatus,
		api.getWifiStatus,
		api.getCloudStatus,
		api.getCloudConfig,
		api.getVoltmeterStatus,
		api.getTemperatureStatus,
		api.getHumidityStatus,
		api.getDevicePowerStatus,
		api.getCoverStatus,
		api.getCoverConfig,
		api.getPM1Status,
		api.getPM1Config,
		api.getEMStatus,
		api.getEMConfig,
		api.getEMDataStatus,
		api.getEM1Status,
		api.getEM1Config,
		api.getEM1DataStatus,
	}

	return api
}
