package main

import (
	"log"
	"testing"

	"github.com/ant0ine/go-json-rest/rest"
)

func TestCallIndex(t *testing.T) {
	app := rest.NewApi()
	app.Use(rest.DefaultDevStack...)
	routes := []*rest.Route{}
	routes = append(routes,
		&rest.Route{
			HttpMethod: "GET",
			PathExp:    "/users",
			Func:       GetUsers,
		},
	)
	router, err := rest.MakeRouter(routes...)
	if err != nil {
		log.Fatal(err)
	}
	app.SetApp(router)
	// recorded := test.RunRequest(t, app.MakeHandler(),
	// makeCustomRequest(ts.Method, "http://1.2.3.4/users", {}, map[string]string{}))
}
