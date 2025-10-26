package repository

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/ryutaKimu/kakebo/internal/model"
	repository "github.com/ryutaKimu/kakebo/internal/repository/user"
)

var _ repository.UserRepository = (*UserRepository)(nil)

type UserRepository struct {
	db   *sql.DB
	goqu goqu.DialectWrapper
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db:   db,
		goqu: goqu.Dialect("postgres"),
	}
}

func (r *UserRepository) CheckUserExists(ctx context.Context, email string) (bool, error) {
	query, args, err := r.goqu.
		From("users").
		Select(goqu.COUNT("id")).
		Where(goqu.I("email").Eq(email)).
		Where(goqu.I("deleted_at").IsNull()).
		ToSQL()
	if err != nil {
		return false, err
	}

	row := r.db.QueryRowContext(ctx, query, args...)
	var count int
	err = row.Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user *model.User) error {
	record := goqu.Record{
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
	}
	query, args, err := r.goqu.Insert("users").Rows(record).ToSQL()

	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}
