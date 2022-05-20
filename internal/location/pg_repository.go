package location

import (
	"context"
)

// Location Repository
type Repository interface {
	FindSatelliteByName(ctx context.Context, name string) (int, error)
}
