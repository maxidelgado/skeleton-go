package example

import "context"

type Service interface {
	Get(context.Context, string) (Entity, error)
}

type Repository interface {
	Get(context.Context, string) (Entity, error)
}
