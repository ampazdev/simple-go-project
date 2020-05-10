package main

import (
	"flag"
	"log"
	"os"

	"github.com/ampazdev/simple-go-project/svc/productservice/internal/app"
	"github.com/ampazdev/simple-go-project/svc/productservice/constant"
	"github.com/ampazdev/simple-go-project/svc/productservice/internal/delivery/rest/echo"
)

func main() {
	env := os.Getenv(constant.OsEnvDefaultEnvironment)
	if env == "" {
		env = constant.OsEnvDefaultEnvironment
	}

	//configPath := flag.String("config", "", "specifies configuration's file path")
	flag.Parse()

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)

	prodServiceApp, err := app.NewProductService(constant.ConfigProjectFilepath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		errs := prodServiceApp.Close()
		for e := range errs {
			log.Println(e)
		}
	}()

	// Start rest server
	echo.Start(prodServiceApp)
}
