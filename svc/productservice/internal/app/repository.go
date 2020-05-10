package app

//import (
//	"fmt"
//	"github.com/ampazdev/simple-go-project/svc/userservice/internal"
//	"github.com/ampazdev/simple-go-project/svc/userservice/internal/app/config"
//	"github.com/ampazdev/simple-go-project/svc/userservice/internal/repo"
//
//	"github.com/ampazdev/simple-go-project/pkg/database/postgresql/connection"
//)
//
//type repository struct {
//	config         config.Config
//	conn           connection.Database
//	userRepository internal.UserRepo
//}
//
//func newRepository(cfg config.Config) (*repository, error) {
//	// var connConfig = connection.Config{DatabaseURI: "userservices=xxx password=xxx dbname=postgres host=localhost port=5433 sslmode=disable"}
//
//	// conn, err := connection.New(&connConfig)
//
//	conn, err := config.NewDatabase(cfg)
//
//	if err != nil {
//		return nil, fmt.Errorf("failed to get Postgre connection ", err)
//	}
//
//	repo := &repository{
//		config: cfg,
//		conn: connection.Database{
//			DB: conn,
//		},
//	}
//
//	repo.userRepository = newUserRepo(
//		connection.Database{
//			DB: conn,
//		})
//
//	return repo, nil
//}
//
//func newUserRepo(conn connection.Database) internal.UserRepo {
//	return repo.NewPqlUserRepository(conn)
//}
//
//func (r *repository) Close() []error {
//	var errs []error
//
//	return errs
//}
