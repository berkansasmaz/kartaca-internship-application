package main

import (
	"github.com/berkansasmaz/kartaca-internship-application.git/api"
	"github.com/berkansasmaz/kartaca-internship-application.git/models"
	"github.com/berkansasmaz/kartaca-internship-application.git/routes"
	"github.com/urfave/negroni"
)

func main() {
	db := models.NewSqliteDB("data.db")
	api := api.NewAPI(db)
	routes := routes.NewRoutes(api)
	n := negroni.Classic()
	n.UseHandler(routes)
	n.Run(":3000")
}
