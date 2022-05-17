package usecase

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"

	"mercadolibre/config"
	"mercadolibre/internal/location"
	"mercadolibre/internal/models"
	"mercadolibre/pkg/logger"
)

const (
	basePrefix    = "api-location:"
	cacheDuration = 3600
)

type locationUC struct {
	cfg          *config.Config
	locationRepo location.Repository
	logger       logger.Logger
}

// Location UseCase constructor
func NewLocationUseCase(cfg *config.Config, locationRepo location.Repository, logger logger.Logger) location.UseCase {
	return &locationUC{cfg: cfg, locationRepo: locationRepo, logger: logger}
}

func (u *locationUC) GetLocationBySatellites(ctx context.Context, satellites models.Request) (*models.Ship, error) {
	u.logger.Error(">>>>>>>>>>>>> usecase http GetLocationBySatellites ini >>>>>>>>>>")
	span, ctx := opentracing.StartSpanFromContext(ctx, "locationUC.GetLocationBySatellites")
	defer span.Finish()
	position := models.Position{X: -100.0, Y: 75.5}
	n := &models.Ship{
		Position: position,
		Message:  "este es un mensaje secreto",
	}

	/*var shipRequest models.Request
	json.NewEncoder(satellites).Encode(&shipRequest)
	newsBase, err := u.redisRepo.GetLocationBySatellitesCtx(ctx, satellites)
	if err != nil {
		u.logger.Errorf("locationUC.GetLocationBySatellites.GetLocationBySatellitesCtx: %v", err)
	}
	if newsBase != nil {
		return newsBase, nil
	}

	n, err := u.locationRepo.GetLocationBySatellites(ctx, satellites)
	if err != nil {
		return nil, err
	}*/
	u.logger.Error(">>>>>>>>>>>>> FIN usecase http GetLocationBySatellites FIN >>>>>>>>>>")
	return n, nil
}

func (u *locationUC) getKeyWithPrefix(newsID string) string {
	return fmt.Sprintf("%s: %s", basePrefix, newsID)
}
