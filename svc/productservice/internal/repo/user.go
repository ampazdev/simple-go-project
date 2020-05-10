package repo

import (
	"context"
	"github.com/ampazdev/simple-go-project/services/userservices/internal/entity"
	"github.com/ampazdev/simple-go-project/pkg/database/postgresql/connection"
)

type pqlUserRepository struct {
	connection.Database
}

func (m *pqlUserRepository) GetByID(ctx context.Context, id int64) (*entity.User, error) {
	// will be added

	return &entity.User{
		ID:   3,
		Name: "Testing",
	}, nil

}

func (m *pqlUserRepository) GetAll(ctx context.Context) (res []*entity.User, err error) {
	query := `SELECT id, name from user_data`

	rows, err := m.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []*entity.User

	for rows.Next() {
		user := new(entity.User)
		err = rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return nil, err
		}

		result = append(result, user)
	}

	return result, nil
}

func NewPqlUserRepository(conn connection.Database) *pqlUserRepository {
	return &pqlUserRepository{Database: conn}
}
