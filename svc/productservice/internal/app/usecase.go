package app

//import (
//	"github.com/ampazdev/simple-go-project/services/userservices/internal"
//	"github.com/ampazdev/simple-go-project/services/userservices/internal/usecase"
//	"time"
//)
//
//type UseCase struct {
//	User internal.UserUseCase
//}
//
//func newUseCase(repository *repository) *UseCase {
//	timeoutContext := time.Duration(30) * time.Second
//	userUsecase := usecase.NewUserUseCase(repository.userRepository, timeoutContext)
//
//	return &UseCase{
//		User: userUsecase,
//	}
//}
//
//func (*UseCase) Close() []error {
//	var errs []error
//	return errs
//}
