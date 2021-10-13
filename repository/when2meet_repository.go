package repository

import (
	"context"

	"github.com/airbenders/when2meet/domain"
	"github.com/airbenders/when2meet/utils/errors"
	"github.com/jackc/pgx/v4/pgxpool"
)

type When2meetRepository struct {
	db *pgxpool.Pool
}

func NewWhen2meetRepository(db *pgxpool.Pool) domain.When2meetRepository {
	return &When2meetRepository{
		db: db,
	}
}

const (
	insert = `INSERT INTO public.When2meet(
	W2W_ID, W2W_name)
	VALUES ($1, $2);`
	selectByID = `SELECT W2W_ID, W2W_name
	FROM public.When2meet WHERE W2W_ID=$1;`
	update = `UPDATE public.When2meet
	SET W2W_name=$2
	WHERE W2W_ID=$1;`
	delete = `DELETE FROM public.When2meet
	WHERE W2W_ID=$1;`
)

func (r *When2meetRepository) Create(ctx context.Context, id string, w2m *domain.When2meet) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, insert, w2m.W2M_ID, w2m.W2M_name)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	err = tx.Commit(ctx)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (r *When2meetRepository) GetByID(ctx context.Context, id string) (*domain.When2meet, error) {
	rows, err := r.db.Query(ctx, selectByID, id)
	if err != nil {
		err = errors.NewInternalServerError(err.Error())
		return nil, err
	}
	defer rows.Close()

	var w2m domain.When2meet
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			err = errors.NewInternalServerError(err.Error())
			return nil, err
		}

		w2m.W2M_ID = values[0].(string)
		w2m.W2M_name = values[1].(string)
	}

	return &w2m, nil
}

func (r *When2meetRepository) Update(ctx context.Context, w2m *domain.When2meet) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, update, w2m.W2M_ID, w2m.W2M_name)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	err = tx.Commit(ctx)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (r *When2meetRepository) Delete(ctx context.Context, id string) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, delete, id)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	err = tx.Commit(ctx)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
