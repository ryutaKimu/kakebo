package repository

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	postgres "github.com/ryutaKimu/kakebo/internal/infra/postgre"
	"github.com/ryutaKimu/kakebo/internal/model"
	repository "github.com/ryutaKimu/kakebo/internal/repository/user"
)

var _ repository.UserRepository = (*UserRepository)(nil)

type dbExecutor interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

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
	exec := getDBExecutor(ctx, r.db)
	query, args, err := r.goqu.
		From("users").
		Select(goqu.COUNT("id")).
		Where(goqu.I("email").Eq(email)).
		ToSQL()
	if err != nil {
		return false, err
	}

	row := exec.QueryRowContext(ctx, query, args...)
	var count int
	err = row.Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user *model.User) error {
	exec := getDBExecutor(ctx, r.db)
	record := goqu.Record{
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
	}
	query, args, err := r.goqu.Insert("users").Rows(record).ToSQL()

	if err != nil {
		return err
	}

	_, err = exec.ExecContext(ctx, query, args...)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) LoginUser(ctx context.Context, email string) (*model.User, error) {
	query, args, err := r.goqu.
		From("users").
		Select("id", "name", "email", "password").
		Where(goqu.I("email").Eq(email)).
		Limit(1).
		ToSQL()
	if err != nil {
		return nil, err
	}

	exec := getDBExecutor(ctx, r.db)
	row := exec.QueryRowContext(ctx, query, args...)
	var user model.User
	err = row.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func getDBExecutor(ctx context.Context, db *sql.DB) dbExecutor {
	if tx, ok := ctx.Value(postgres.TxContextKey).(*sql.Tx); ok && tx != nil {
		return tx
	}
	return db
}
