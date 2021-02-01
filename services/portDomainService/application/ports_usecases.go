package application

import (
	"context"
	"errors"
	"fmt"
	"ndanamedtt/services/portDomainService/domain"
)

var ErrNotFound = errors.New("entity not found")

type PortFinderFunc func(finder PortFinder) PortFinder

type PortFinder interface {
	WithID(domain.PortID) PortFinder
	WithName(string) PortFinder
	WithCity(string) PortFinder
	WithCountry(string) PortFinder
	WithLimit(int) PortFinder
	Find(ctx context.Context) ([]domain.Port, error)
	FindOne(ctx context.Context) (domain.Port, error)
}

type CreatePortParams struct {
	Name        string
	City        string
	Country     string
	Alias       []string
	Regions     []string
	Coordinates []float64
	Province    string
	Timezone    string
	Unlocs      []string
	Code        string
}

type PortRepository interface {
	Store(context.Context, CreatePortParams) (domain.Port, error)
	Update(context.Context, domain.Port) error
}

type CreateOrUpdatePortFunc func(context.Context, CreatePortParams) error

func NewCreateOrUpdatePort(pf PortFinder, pr PortRepository) CreateOrUpdatePortFunc {
	return func(ctx context.Context, params CreatePortParams) error {
		handleErr := func(err error) error {
			return fmt.Errorf("CreateOrUpdaterPort: %w", err)
		}
		oldP, err := pf.WithName(params.Name).WithCity(params.City).WithCountry(params.Country).FindOne(ctx)
		if err != nil && !errors.Is(err, ErrNotFound) {
			return handleErr(err)
		}
		if errors.Is(err, ErrNotFound) {
			_, err := pr.Store(ctx, params)
			if err != nil {
				return handleErr(err)
			}
		}
		updatedPort := domain.Port{
			ID:          oldP.ID,
			Name:        params.Name,
			City:        params.City,
			Country:     params.Country,
			Alias:       params.Alias,
			Regions:     params.Regions,
			Coordinates: params.Coordinates,
			Province:    params.Province,
			Timezone:    params.Timezone,
			Unlocs:      params.Unlocs,
			Code:        params.Country,
		}
		if err := pr.Update(ctx, updatedPort); err != nil {
			return handleErr(err)
		}
		return nil
	}
}
