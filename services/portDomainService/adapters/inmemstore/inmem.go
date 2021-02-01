package inmemstore

import (
	"context"
	"ndanamedtt/services/portDomainService/application"
	"ndanamedtt/services/portDomainService/domain"
)

type InMemRepo struct{}

func (InMemRepo) Store(ctx context.Context, params application.CreatePortParams) (domain.Port, error) {
	panic("implement me")
}

func (InMemRepo) Update(ctx context.Context, port domain.Port) error {
	panic("implement me")
}
