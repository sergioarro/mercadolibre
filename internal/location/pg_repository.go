package location

import (
	"context"
	"mercadolibre/internal/models"
	"sync"
)

// Location Repository
type Repository interface {
	FindSatelliteByName(ctx context.Context, name string) (int, error)
	Create(ctx context.Context, satellite models.Satellite, wg *sync.WaitGroup) error
	GetAllSatellites(ctx context.Context, wg *sync.WaitGroup) ([]models.Satellite, error)
}
