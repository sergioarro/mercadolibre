package repository

import (
	"context"
	"encoding/json"
	"log"
	"sync"

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

type PositionMap map[string]interface{}

// Location repository constructor
func NewLocationRepository(db *sqlx.DB) location.Repository {
	return &locationRepo{db: db}
}

/*
func (p PositionMap) Value() (driver.Value, error) {
	j, err := json.Marshal(p)
	return j, err
}

func (p *PositionMap) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	var i interface{}
	err := json.Unmarshal(source, &i)
	if err != nil {
		return err
	}

	*p, ok = i.(map[string]interface{})
	if !ok {
		return errors.New("Type assertion .(map[string]interface{}) failed.")
	}

	return nil
}
*/

func (r *locationRepo) FindSatelliteByName(ctx context.Context, name string) (int, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "locationRepo.FindSatelliteByName")
	defer span.Finish()

	var totalCount int
	if err := r.db.GetContext(ctx, &totalCount, getSatelliteByNameQuery, name); err != nil {
		return 0, errors.Wrap(err, "locationRepo.FindSatelliteByName.GetContext")
	}

	return totalCount, nil
}

func (r *locationRepo) GetAllSatellites(ctx context.Context, wg *sync.WaitGroup) ([]models.Satellite, error) {
	defer wg.Done()
	span, ctx := opentracing.StartSpanFromContext(ctx, "locationRepo.GetAllSatellites")
	defer span.Finish()

	var satellites = make([]models.Satellite, 0, 2)

	rows, err := r.db.QueryxContext(ctx, getSatellites)

	defer rows.Close()

	for rows.Next() {
		s := models.Satellite{}

		if err = rows.Scan(&s.Name, (*pq.StringArray)(&s.Message), &s.Distance, &s.Position); err != nil {
			return nil, errors.Wrap(err, "locationRepo.GetAllSatellites.StructScan")
		}

		log.Printf("----------")
		log.Printf("%+v", s)
		log.Printf("----------")

		satellites = append(satellites, s)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "locationRepo.GetAllSatellites.rows.Err")
	}

	return satellites, nil
}

func (r *locationRepo) Create(ctx context.Context, satellite models.Satellite, wg *sync.WaitGroup) error {
	defer wg.Done()
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
