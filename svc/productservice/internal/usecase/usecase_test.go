package usecase

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/ampazdev/simple-go-project/services/userservices/internal"
	"github.com/ampazdev/simple-go-project/services/userservices/internal/entity"
	mock_internal "github.com/ampazdev/simple-go-project/services/userservices/internal/mocks"
	"reflect"
	"testing"
	"time"
)

func TestNewUserUseCase(t *testing.T) {
	type args struct {
		userRepo       internal.Repository
		contextTimeout time.Duration
	}
	tests := []struct {
		name string
		args args
		want *userUseCase
	}{
		{
			name: "Test new userservices usecase",
			args: args{
				userRepo:       nil,
				contextTimeout: 0,
			},
			want: &userUseCase{
				userRepo:       nil,
				contextTimeout: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserUseCase(tt.args.userRepo, tt.args.contextTimeout); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserUseCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUseCase_GetByID(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockRepo := mock_internal.NewMockRepository(ctl)

	type fields struct {
		userRepo       internal.Repository
		contextTimeout time.Duration
	}
	type args struct {
		c  context.Context
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.User
		wantErr bool
		err     error
	}{
		{
			name: "Test get userservices by id (success)",
			fields: fields{
				userRepo:       mockRepo,
				contextTimeout: 0,
			},
			args: args{
				c:  context.TODO(),
				id: int64(0),
			},
			want: &entity.User{
				ID:   3,
				Name: "Test",
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Test get userservices by id (error)",
			fields: fields{
				userRepo:       mockRepo,
				contextTimeout: 0,
			},
			args: args{
				c:  context.TODO(),
				id: int64(0),
			},
			want:    nil,
			wantErr: true,
			err:     errors.New("error cuy"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockRepo.EXPECT().GetByID(gomock.Any(), tt.args.id).Return(tt.want, tt.err)

			u := userUseCase{
				userRepo:       mockRepo,
				contextTimeout: tt.fields.contextTimeout,
			}
			got, err := u.GetByID(tt.args.c, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userUseCase_GetAll(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockRepo := mock_internal.NewMockRepository(ctl)

	type fields struct {
		userRepo       internal.Repository
		contextTimeout time.Duration
	}
	type args struct {
		c context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*entity.User
		wantErr bool
		err     error
	}{
		{
			name: "Test Get All User (Success)",
			fields: fields{
				userRepo:       mockRepo,
				contextTimeout: 0,
			},
			args: args{
				c: context.TODO(),
			},
			want: []*entity.User{
				{
					ID:   0,
					Name: "Wikan",
				},
				{
					ID:   2,
					Name: "Haryo",
				},
			},
			wantErr: false,
			err:     nil,
		},
		{
			name: "Test Get All User (Error)",
			fields: fields{
				userRepo:       mockRepo,
				contextTimeout: 0,
			},
			args: args{
				c: context.TODO(),
			},
			want: nil,
			wantErr: true,
			err:     errors.New("ga da data oik"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockRepo.EXPECT().GetAll(gomock.Any()).Return(tt.want, tt.err)

			u := userUseCase{
				userRepo:       mockRepo,
				contextTimeout: tt.fields.contextTimeout,
			}
			got, err := u.GetAll(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}
