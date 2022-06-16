package service

import (
	"context"
	"grsidecar/pbgen"

	log "github.com/sirupsen/logrus"
)

type AlertRequestServer struct {
	pbgen.UnimplementedCaterAlertRequestServer
	AlertsMem *AlertStore
}

func NewAlertCaterServer(AS *AlertStore) *AlertRequestServer { //if assoicated with AlertStore struct, mutex is passed by value
	AS.ReadyAlertStore()
	return &AlertRequestServer{AlertsMem: AS}
}

func (s *AlertRequestServer) CaterAlert(ctx context.Context, req *pbgen.AlertRequest,
) (*pbgen.AlertResponse, error) {

	alert := req.GetAlertsbatch()
	if alert == nil {
		log.Error("Sorry, you need to provide the parameters! Please call DescribeAlert for info!")
		return &pbgen.AlertResponse{Seen: "No: No alert was recorded"}, nil
	} else {
		log.Infof("Received request to service an alert with name : %s", req.Cid)
		s.AlertsMem.StoreAlert(alert)
		log.Debug("Alert is Stored : %v", s.AlertsMem.store)
	}
	//pass to relevant processes to handle
	return &pbgen.AlertResponse{Seen: "Alert will be serviced soon"}, nil
}
