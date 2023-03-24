package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SiddhantGupta07/bookings/pkg/config"
	"github.com/SiddhantGupta07/bookings/pkg/handlers"
	"github.com/SiddhantGupta07/bookings/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main function
func main() {

	//change this to true when in production

	app.InProduction = false

	session = scs.New()

	session.Lifetime = 24 * time.Hour

	session.Cookie.Persist = true

	session.Cookie.SameSite = http.SameSiteLaxMode

	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	//Creating the http server
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	//Listening and serving the server
	err = srv.ListenAndServe()
	log.Fatal(err)
}
