package example

import "context"

func New(r Repository) Service {
	return service{r: r}
}

type service struct {
	r Repository
}

func (s service) Get(ctx context.Context, id string) (Entity, error) {
	return s.r.Get(ctx, id)
}
