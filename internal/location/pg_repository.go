package location

import (
	"context"

	"mercadolibre/internal/models"
)

// Location Repository
type Repository interface {
	GetLocationBySatellites(ctx context.Context, location models.Request) (*models.Response, error)
}
