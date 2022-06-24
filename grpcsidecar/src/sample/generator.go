package sample

import (
	"grsidecar/pbgen"

	"github.com/google/uuid"
)

func NewAlert() *pbgen.AlertRequest {
	return &pbgen.AlertRequest{
		Cid:         uuid.New().String(),
		Alertsbatch: NewAlertsList(),
	}
}

//TODO: reandomize with for loop initialization later
func NewAlertsList() *pbgen.AlertList {
	return &pbgen.AlertList{
		Alerts: []*pbgen.AlertDetail{randomAlert(), randomAlert(), randomAlert()},
	}
}
