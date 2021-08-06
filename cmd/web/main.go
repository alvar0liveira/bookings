package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/alvaroliveira20/bookings/pkg/config"
	"github.com/alvaroliveira20/bookings/pkg/handlers"
	"github.com/alvaroliveira20/bookings/pkg/render"
)

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()

	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc

	render.NewTemplates(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
