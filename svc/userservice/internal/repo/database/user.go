package database

import (
	"context"
	"github.com/ampazdev/simple-go-project/svc/userservice/internal/entity"
)

func (p *UserRepo) GetByUserInfoByEmail(ctx context.Context, user entity.User) (*entity.User, error) {
	stmt, err := p.DB.PrepareContext(ctx, getUserInfoByEmail)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRowContext(ctx, &user.Email)

	res := entity.User{}
	err = row.Scan(
		&res.ID,
		&res.Email,
		&res.FullName,
		&res.PhoneNumber,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (p *UserRepo) InsertUserInfo(ctx context.Context, user entity.User) (*entity.User, error) {
	stmt, err := p.DB.PrepareContext(ctx, insertUserInfo)
	if err != nil {
		return nil, err
	}

	row := stmt.QueryRowContext(ctx, &user.ID, &user.Email, &user.FullName, &user.PhoneNumber)

	res := entity.User{}
	err = row.Scan(
		&res.ID,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
