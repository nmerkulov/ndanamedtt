package grpc

import (
	"ndanamedtt/pkg/grpc/portDomain"
	"ndanamedtt/services/portDomainService/application"
)

func createPortParamsToApp(p *portDomain.Port) application.CreatePortParams {
	return application.CreatePortParams{
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
