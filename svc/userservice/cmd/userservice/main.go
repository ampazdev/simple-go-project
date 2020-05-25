package main

import (
	"flag"
	"log"
	"os"

	"github.com/ampazdev/simple-go-project/svc/userservice/constant"
	"github.com/ampazdev/simple-go-project/svc/userservice/internal/app"
	"github.com/ampazdev/simple-go-project/svc/userservice/internal/delivery/rest/echo"
)

func main() {
	env := os.Getenv(constant.OsEnvDefaultEnvironment)
	if env == "" {
		env = constant.OsEnvDefaultEnvironment
	}

	//configPath := flag.String("config", "", "specifies configuration's file path")
	flag.Parse()

	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)

	userServiceApp, err := app.NewUserservice(constant.ConfigProjectFilepath)
	if err != nil {
		log.Fatal(err)
	}
	os.Setenv("SECRET", constant.SecretJWT)
	defer func() {
		errs := userServiceApp.Close()
		for e := range errs {
			log.Println(e)
		}
	}()

	// Start rest server
	echo.Start(userServiceApp)
}
