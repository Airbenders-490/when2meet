package repository

import (
	"context"

	"github.com/airbenders/when2meet/domain"
	"github.com/airbenders/when2meet/utils/errors"
	"github.com/jackc/pgx/v4/pgxpool"
)

type AvailibilitiesRepository struct {
	db *pgxpool.Pool
}

func NewAvailibilitiesRepository(db *pgxpool.Pool) domain.AvailibilitiesRepository {
	return &AvailibilitiesRepository{
		db: db,
	}
}

// const (
// 	insert = `INSERT INTO public.availibilities(AID, SID, ADAY, TIMES)
// 	VALUES ($1, $2, $3, $4);`
// 	selectByID = `SELECT W2W_ID, W2W_name
// 	FROM public.Availibilities WHERE W2W_ID=$1;`
// 	update = `UPDATE public.Availibilities
// 	SET W2W_name=$2
// 	WHERE W2W_ID=$1;`
// 	delete = `DELETE FROM public.Availibilities
// 	WHERE W2W_ID=$1;`
// )

func (r *AvailibilitiesRepository) Create(ctx context.Context, id string, ava *domain.Availibilities) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, insert, ava.AID, ava.SID, ava.ADAY, ava.Times)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	err = tx.Commit(ctx)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (r *AvailibilitiesRepository) GetByID(ctx context.Context, id string) (*domain.Availibilities, error) {
	rows, err := r.db.Query(ctx, selectByID, id)
	if err != nil {
		err = errors.NewInternalServerError(err.Error())
		return nil, err
	}
	defer rows.Close()

	var ava domain.Availibilities
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			err = errors.NewInternalServerError(err.Error())
			return nil, err
		}

		ava.AID = values[0].(string)
		ava.SID = values[1].(string)
		ava.ADAY = values[2].(string)
		ava.Times = values[3].([2]int)

	}

	return &ava, nil
}

func (r *AvailibilitiesRepository) Update(ctx context.Context, ava *domain.Availibilities) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, update, ava.AID, ava.SID, ava.ADAY, ava.Times)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	err = tx.Commit(ctx)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (r *AvailibilitiesRepository) Delete(ctx context.Context, id string) error {
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
