package database

import (
	"context"
	"database/sql"
	"github.com/gomodule/redigo/redis"

	"github.com/ampazdev/simple-go-project/svc/userservice/internal"
	"github.com/ampazdev/simple-go-project/svc/userservice/internal/entity"
)

type UserRepo struct {
	DB    *sql.DB // TODO: create bridges
	Cache *redis.Pool
}

func NewUserReaderRepo(db *sql.DB, cache *redis.Pool) internal.UserReaderRepo {
	return &UserRepo{
		DB:    db,
		Cache: cache,
	}
}

func (p *UserRepo) GetUserCredByEmail(ctx context.Context, user entity.User) (*entity.User, error) {
	stmt, err := p.DB.PrepareContext(ctx, getUserCredByEmail)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRowContext(ctx, &user.Email)

	res := entity.User{}
	err = row.Scan(
		&res.ID,
		&res.Email,
		&res.Password,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (p *UserRepo) InsertUserCred(ctx context.Context, user entity.User) (*entity.User, error) {
	stmt, err := p.DB.PrepareContext(ctx, insertUserLoginCred)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRowContext(ctx, &user.Email, &user.Password)

	res := entity.User{}
	err = row.Scan(
		&res.ID,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (p *UserRepo) SignupUser(ctx context.Context, user entity.User) (*entity.User, error) {
	return &entity.User{}, nil
}

func (p *UserRepo) Login(ctx context.Context, user entity.User) (*entity.User, error) {
	return &entity.User{}, nil
}
