 package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/rdturbo/bookings/pkg/config"
	"github.com/rdturbo/bookings/pkg/handlers"
	"github.com/rdturbo/bookings/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {
	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 *time.Hour
	session.Cookie.Persist = true // for session to persist after browser is closed
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // for https (always true in production mode)

	app.Session = session // so that handlers can access this session variable through config file

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("Starting application on the port %s", portNumber)
	
	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
