package dataaccess

import (
	"context"

	"github.com/maxidelgado/skeleton-go/domain/example"
)

func New() example.Repository {
	return datastore{}
}

type datastore struct {
}

func (d datastore) Get(ctx context.Context, id string) (example.Entity, error) {
	return example.Entity{Id: id}, nil
}
