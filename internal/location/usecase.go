package location

import (
	"context"

	"mercadolibre/internal/models"
)

type UseCase interface {
	GetLocationBySatellites(ctx context.Context, location models.Request) (*models.Ship, error)
}
