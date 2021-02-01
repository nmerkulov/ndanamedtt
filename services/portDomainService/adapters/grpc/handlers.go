package grpc

import (
	"context"
	"errors"
	"fmt"
	"io"
	"ndanamedtt/pkg/grpc/portDomain"
	"ndanamedtt/services/portDomainService/application"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PortsService struct {
	portDomain.UnimplementedPortDomainServer
	createOrUpdatePort application.CreateOrUpdatePortFunc
}

func (ps PortsService) RecordPort(server portDomain.PortDomain_RecordPortServer) error {
	stats := portDomain.RecordPortStats{}
	for {
		port, err := server.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return status.Error(codes.Internal, err.Error())
		}
		stats.Total++
		if err := ps.createOrUpdatePort(server.Context(), createPortParamsToApp(port)); err != nil {
			stats.Rejected++
			continue
		}
		stats.Accepted++
	}
	if err := server.SendAndClose(&stats); err != nil {
		return fmt.Errorf("PortsService.RecordPort")
	}
	return nil
}

func (ps PortsService) FindPort(ctx context.Context, params *portDomain.FindPortParams) (*portDomain.Port, error) {
	panic("implement me")
}
