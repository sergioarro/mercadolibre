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

// FindSatelliteByName
func (r *locationRepo) FindSatelliteByName(ctx context.Context, name string) (*models.Satellite, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "locationRepo.FindSatelliteByName")
	defer span.Finish()

	satellite := &models.Satellite{}
	if err := r.db.QueryRowxContext(ctx, getSatelliteByNameQuery, name).StructScan(satellite); err != nil {
		return nil, errors.Wrap(err, "locationRepo.FindSatelliteByName.QueryRowxContext")
	}

	return satellite, nil
}
