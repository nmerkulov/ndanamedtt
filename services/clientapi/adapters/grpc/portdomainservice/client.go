package portdomainservice

import (
	"context"
	"fmt"
	"ndanamedtt/pkg/grpc/portDomain"
	"ndanamedtt/services/clientapi/application"
	"ndanamedtt/services/clientapi/domain"

	"google.golang.org/grpc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcPortDomainClient struct {
	c portDomain.PortDomainClient
}

func NewGrpcPortDomainClient(conn *grpc.ClientConn) GrpcPortDomainClient {
	return GrpcPortDomainClient{
		c: portDomain.NewPortDomainClient(conn),
	}
}

func (g GrpcPortDomainClient) StreamPorts(ctx context.Context, ports <-chan domain.Port) (application.StreamResult, error) {
	handleErr := func(err error) (application.StreamResult, error) {
		return application.StreamResult{}, fmt.Errorf("GrpcPortDomainClient.StreamPorts: %w", err)
	}
	s, err := g.c.RecordPort(ctx)
	if err != nil {
		return handleErr(err)
	}
	for p := range ports {
		if err := s.Send(portToGrpc(p)); err != nil {
			return handleErr(err)
		}
	}
	resp, err := s.CloseAndRecv()
	if err != nil {
		return handleErr(err)
	}
	return application.StreamResult{
		Total:    resp.Total,
		Accepted: resp.Accepted,
		Rejected: resp.Rejected,
	}, nil
}

func (g GrpcPortDomainClient) FindPortByID(ctx context.Context, id int64) (domain.Port, error) {
	p, err := g.c.FindPort(ctx, &portDomain.FindPortParams{Id: id})
	if err != nil {
		if status.Convert(err).Code() == codes.NotFound {
			return domain.Port{}, application.ErrNotFound
		}
		return domain.Port{}, fmt.Errorf("GrpcPortDomainClient.FindPortByName: %w", err)
	}
	return portToDomain(p), nil
}
