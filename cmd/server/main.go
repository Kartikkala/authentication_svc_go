package main

import (
	"fmt"
	"log"

	"github.com/Kartikkala/authentication_svc/authentication"
	"github.com/Kartikkala/authentication_svc/config"
	"github.com/Kartikkala/authentication_svc/hooks"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
)

func main() {
	app, err := config.NewApp()
	if err != nil {
		fmt.Println(err.Error())
	}

	nc, err := nats.Connect(app.Cfg.NATS.URL)

	if err != nil {
		log.Println("Error in NATS server connection...", err)
		return
	}

	authenticationSvc := authentication.NewService(app.DB, *app.Cfg)

	authenticationSvcHooks := hooks.NewAuthenticationSvcHooks(nc)

	authenticationSvcWithHooks := authentication.NewServiceWithHooks(authenticationSvc)
	authenticationSvcWithHooks.AppendAfterRegistrationHook(authenticationSvcHooks.NotifyDriveService)

	
	e := echo.New()

	// Returns decode token middleware, maybe could be used
	// in future for refresh tokens
	authentication.AttachRoutes(e, authenticationSvcWithHooks)
	
	port := app.Cfg.App.RESTPort
	address := app.Cfg.App.HostAddress
	log.Printf("Starting HTTP service on %v:%d", address, port)
	e.Start(fmt.Sprintf("%s:%d", address, port))
}
