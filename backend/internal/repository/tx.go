package repository

import (
	"context"
)

type dbtx interface {
	SelectContext(ctx context.Context, dest any, query string, args ...any) error
}
