package portdomainservice

import (
	"ndanamedtt/pkg/grpc/portDomain"
	"ndanamedtt/services/clientapi/domain"
)

func portToGrpc(p domain.Port) *portDomain.Port {
	return &portDomain.Port{
		Name:        p.Name,
		City:        p.City,
		Country:     p.Country,
		Alias:       p.Alias,
		Regions:     p.Regions,
		Coordinates: p.Coordinates,
		Province:    p.Province,
		Timezone:    p.Timezone,
		Unlocs:      p.Unlocs,
		Code:        p.Code,
	}
}

func portToDomain(p *portDomain.Port) domain.Port {
	return domain.Port{
		Name:        p.Name,
		City:        p.City,
		Country:     p.Country,
		Alias:       p.Alias,
		Regions:     p.Regions,
		Coordinates: p.Coordinates,
		Province:    p.Province,
		Timezone:    p.Timezone,
		Unlocs:      p.Unlocs,
		Code:        p.Code,
	}
}
