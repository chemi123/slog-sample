package infra_interface

import "context"

type Repository interface {
	Get(ctx context.Context) int
}
