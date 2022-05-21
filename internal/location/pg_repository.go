package location

import (
	"context"
	"mercadolibre/internal/models"
)

// Location Repository
type Repository interface {
	FindSatelliteByName(ctx context.Context, name string) (int, error)
	Create(ctx context.Context, satellite models.Satellite) error
}
