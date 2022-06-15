package service

import (
	"context"
	"grsidecar/pbgen"
)

type AlertRequestServer struct {
	pbgen.UnimplementedSendAlertRequestServer
	AlertStorage []string
}

func (s *AlertRequestServer) SendAlert(ctx context.Context, req *pbgen.AlertRequest,
) (*pbgen.AlertResponse, error) {

	return &pbgen.AlertResponse{}, nil
}
