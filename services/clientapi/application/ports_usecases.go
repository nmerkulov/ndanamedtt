package application

import (
	"context"
	"errors"
	"fmt"
	"io"
	"ndanamedtt/services/clientapi/domain"
)

var ErrNotFound = errors.New("entity not found")

type StreamResult struct {
	Total    int64
	Accepted int64
	Rejected int64
}

type PortsSourceFunc func(ctx context.Context) <-chan domain.Port
type FindPortsProxyFunc func(context.Context, int64) (domain.Port, error)

type PortDomainService interface {
	FindPortByID(context.Context, int64) (domain.Port, error)
	StreamPorts(context.Context, <-chan domain.Port) (StreamResult, error)
}

func (s StreamResult) Print(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "Total: %d\nAccepted: %d\nRejected: %d/n", s.Total, s.Accepted, s.Rejected); err != nil {
		return fmt.Errorf("StreamResult.Print: %w", err)
	}
	return nil
}

func ConsumePorts(sf PortsSourceFunc, ps PortDomainService, w io.Writer) func(context.Context) error {
	return func(ctx context.Context) error {
		handleErr := func(err error) error {
			return fmt.Errorf("ConsumePorts: %w", err)
		}
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		ports := sf(ctx)
		stats, err := ps.StreamPorts(ctx, ports)
		if err != nil {
			return handleErr(err)
		}
		if err := stats.Print(w); err != nil {
			return fmt.Errorf("ConsumePorts: %w", err)
		}
		return nil
	}
}

func FindPortsProxy(ps PortDomainService) FindPortsProxyFunc {
	return func(ctx context.Context, ID int64) (domain.Port, error) {
		p, err := ps.FindPortByID(ctx, ID)
		if err != nil {
			return domain.Port{}, fmt.Errorf("FindPorts: %w", err)
		}
		return p, nil
	}
}
