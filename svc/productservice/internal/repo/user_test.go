package repo

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/isfanazha/simple-go-project/pkg/database/postgresql/connection"
	"github.com/isfanazha/simple-go-project/services/userservices/internal/entity"
	"github.com/jmoiron/sqlx"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestNewPqlUserRepository(t *testing.T) {
	db, _, err := sqlmock.New()

	dbMock := connection.Database{&sqlx.DB{
		DB: db,
	}}

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	type args struct {
		conn connection.Database
	}
	tests := []struct {
		name string
		args args
		want *pqlUserRepository
	}{
		{
			name: "Testing constructor",
			args: args{
				conn: dbMock,
			},
			want: &pqlUserRepository{
				dbMock,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPqlUserRepository(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPqlUserRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pqlUserRepository_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	dbMock := connection.Database{&sqlx.DB{
		DB: db,
	}}

	type fields struct {
		Database connection.Database
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantRes []*entity.User
		wantErr bool
		rowErr  bool
		err     error
	}{
		{
			name: "test get all (success)",
			fields: fields{
				Database: dbMock,
			},
			args: args{
				ctx: context.TODO(),
			},
			wantRes: []*entity.User{
				{
					ID:   0,
					Name: "Wikan",
				},
				{
					ID:   1,
					Name: "Haryo",
				},
			},
			wantErr: false,
			rowErr:  false,
			err:     nil,
		},
		{
			name: "test get all (failed)",
			fields: fields{
				Database: dbMock,
			},
			args: args{
				ctx: context.TODO(),
			},
			wantRes: nil,
			wantErr: true,
			rowErr:  false,
			err:     errors.New("error"),
		},
		{
			name: "test get all (failed while parsing)",
			fields: fields{
				Database: dbMock,
			},
			args: args{
				ctx: context.TODO(),
			},
			wantRes: nil,
			wantErr: true,
			rowErr:  true,
			err:     errors.New("error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := "SELECT id, name from user_data"

			if tt.wantRes != nil {

				rows := sqlmock.NewRows([]string{"id", "name"}).
					AddRow(tt.wantRes[0].ID, tt.wantRes[0].Name).
					AddRow(tt.wantRes[1].ID, tt.wantRes[1].Name)

				mock.ExpectQuery(query).WillReturnRows(rows)
			}

			if tt.rowErr {
				rows := sqlmock.NewRows([]string{"id", "name"}).
					AddRow("oke", "wkwk").
					AddRow(1, "zz")

				mock.ExpectQuery(query).WillReturnRows(rows)
			}

			m := &pqlUserRepository{
				Database: tt.fields.Database,
			}

			gotRes, err := m.GetAll(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("GetAll() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
