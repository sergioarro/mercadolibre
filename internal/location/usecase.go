package location

import (
	"context"

	"mercadolibre/internal/models"
)

type UseCase interface {
	GetLocationBySatellites(ctx context.Context, location models.Request) (*models.Response, error)
	PostTopSecretSplit(ctx context.Context, satelliteName string, location models.Satellite) (*models.Response, error)
	GetTopSecretSplit(ctx context.Context) (*models.Response, error)
}
