package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"mercadolibre/internal/location"
	"mercadolibre/internal/models"
)

// Location Repository
type locationRepo struct {
	db *sqlx.DB
}

// Location repository constructor
func NewLocationRepository(db *sqlx.DB) location.Repository {
	return &locationRepo{db: db}
}

// Get single location by id
func (r *locationRepo) GetLocationBySatellites(ctx context.Context, location models.Request) (*models.Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "locationRepo.GetLocationByID")
	defer span.Finish()

	n := &models.Response{}
	if err := r.db.GetContext(ctx, n, getLocationBySatellites, location); err != nil {
		return nil, errors.Wrap(err, "locationRepo.GetLocationByID.GetContext")
	}

	return n, nil
}
