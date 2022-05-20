package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"

	"github.com/pkg/errors"

	"mercadolibre/internal/location"
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
func (r *locationRepo) FindSatelliteByName(ctx context.Context, name string) (int, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "locationRepo.FindSatelliteByName")
	defer span.Finish()

	var totalCount int
	if err := r.db.GetContext(ctx, &totalCount, getSatelliteByNameQuery, name); err != nil {
		return 0, errors.Wrap(err, "locationRepo.FindSatelliteByName.GetContext")
	}

	return totalCount, nil
}
