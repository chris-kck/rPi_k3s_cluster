package timescaledb

import (
	"context"

	"github.com/sss-eda/lemi-011b/pkg/domain/acquisition"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Repository TODO
type Repository struct {
	pool *pgxpool.Pool
}

// NewRepository TODO
func NewRepository(
	ctx context.Context,
	pgxpool *pgxpool.Pool,
) (*Repository, error) {
	repo := &Repository{
		pool: pgxpool,
	}

	_, err := repo.pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS sensor
		(
			id SERIAL PRIMARY KEY
		);

		CREATE TABLE IF NOT EXISTS datum
		(
			time TIMESTAMPTZ PRIMARY KEY,
			sensor_id INTEGER,
			x INTEGER,
			y INTEGER,
			z INTEGER,
			t INTEGER,
			FOREIGN KEY (sensor_id) REFERENCES sensor (id)
		);

		SELECT create_hypertable('datum', 'time', if_not_exists => TRUE);
	`)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

// AcquireDatum TODO
func (repo *Repository) AcquireDatum(
	ctx context.Context,
	datum acquisition.Datum,
) error {
	_, err := repo.pool.Exec(ctx, `
		INSERT INTO datum (time, sensor_id, x, y, z, t)
		VALUES ($1, $2, $3, $4, $5, $6);
	`, datum.Time, datum.SensorID, datum.X, datum.Y, datum.Z, datum.T)
	if err != nil {
		return err
	}

	return nil
}
