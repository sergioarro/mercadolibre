package repository

import (
	"context"
	"encoding/json"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
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
func (r *locationRepo) FindSatelliteByName(ctx context.Context, name string) (int, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "locationRepo.FindSatelliteByName")
	defer span.Finish()

	var totalCount int
	if err := r.db.GetContext(ctx, &totalCount, getSatelliteByNameQuery, name); err != nil {
		return 0, errors.Wrap(err, "locationRepo.FindSatelliteByName.GetContext")
	}

	return totalCount, nil
}

func (r *locationRepo) Create(ctx context.Context, satellite models.Satellite) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "locationRepo.Create")
	defer span.Finish()

	c := &models.Satellite{}
	p, err := json.Marshal(&satellite.Position)
	if err = r.db.QueryRowxContext(
		ctx,
		createSatellite,
		&satellite.Name,
		pq.Array(&satellite.Message),
		&satellite.Distance,
		p,
	).StructScan(c); err != nil {
		return errors.Wrap(err, "locationRepo.Create.StructScan")
	}

	return nil
}
