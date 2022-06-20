package sample

import "grsidecar/pbgen"

func NewAlert() *pbgen.AlertRequest {
	return &pbgen.AlertRequest{
		Cid:         "",
		Alertsbatch: NewAlertsList(),
	}
}

//TODO: reandomize with for loop initialization later
func NewAlertsList() *pbgen.AlertList {
	return &pbgen.AlertList{
		Alerts: []*pbgen.AlertDetail{randomAlert(), randomAlert(), randomAlert()},
	}
}
