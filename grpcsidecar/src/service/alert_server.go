package service

import (
	"context"
	"grsidecar/pbgen"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		return nil, status.Errorf(codes.InvalidArgument,
			"Sorry, you need to provide the parameters! Please call UnimplementedDescribeAlert for info!")
	} else {
		log.Infof("Received request to service an alert with cid : %s", req.Cid)
		s.AlertsMem.StoreAlert(alert)
		log.Debug("Alert is Stored : %v", s.AlertsMem.store)
	}

	if ctx.Err() == context.DeadlineExceeded {
		log.Printf("Deadline exceeded!")
		return nil, status.Error(codes.DeadlineExceeded, "Deadline exceeded")
	}

	if ctx.Err() == context.Canceled {
		log.Printf("Request was canceled!")
		return nil, status.Error(codes.Canceled, "Context canceled")
	}
	//pass to relevant processes to handle
	return &pbgen.AlertResponse{Seen: "Alert will be serviced soon"}, nil
}
