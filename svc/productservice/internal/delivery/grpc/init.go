package grpc

import (
	"log"

	productproto "github.com/ampazdev/simple-go-project/myproto/productservice"
	"github.com/ampazdev/simple-go-project/svc/productservice/internal/app"
	"github.com/ampazdev/simple-go-project/svc/productservice/internal/delivery/grpc/service"
)

func Start(app *app.ProductService) {
	// get all services
	svc := service.GetServices(app)

	// create new GRPC server
	server := NewServer(app.Config.GRPC.Port)
	// register server
	productproto.RegisterProductServiceServer(server.GRPCServer, svc)

	log.Printf("GRPC server is starting on %s", app.Config.GRPC.Port)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
