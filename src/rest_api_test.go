package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/ant0ine/go-json-rest/rest/test"
)

type TestSimpleApi struct {
	Url     string            //url router
	Func    rest.HandlerFunc  //handler function for api
	Method  string            //GET,POST,PUT,DELETE
	ReqUrl  string            //url client request
	Payload interface{}       //Payload of request
	Header  map[string]string //Header of request
}

func (ts *TestSimpleApi) RunRequest(t *testing.T) *test.Recorded {
	app := rest.NewApi()
	app.Use(rest.DefaultDevStack...)
	routes := []*rest.Route{}
	routes = append(routes,
		&rest.Route{
			HttpMethod: ts.Method,
			PathExp:    ts.Url,
			Func:       ts.Func,
		},
	)
	router, err := rest.MakeRouter(routes...)
	if err != nil {
		log.Fatal(err)
	}
	app.SetApp(router)

	recorded := test.RunRequest(t, app.MakeHandler(),
		makeCustomRequest(ts.Method, "http://1.2.3.4"+ts.ReqUrl, ts.Payload, ts.Header))
	return recorded
}

func makeCustomRequest(method string, urlStr string, payload interface{}, header map[string]string) *http.Request {
	var s string

	if payload != nil {
		b, err := json.Marshal(payload)
		if err != nil {
			panic(err)
		}
		s = fmt.Sprintf("%s", b)
	}

	r, err := http.NewRequest(method, urlStr, strings.NewReader(s))
	if err != nil {
		panic(err)
	}
	r.Header.Set("Accept-Encoding", "gzip")

	for k, v := range header {
		r.Header.Set(k, v)
	}
	if payload != nil {
		r.Header.Set("Content-Type", "application/json")
	}

	return r
}

func TestCallGetUsers(t *testing.T) {
	users := Users{
		Store: map[string]*User{},
	}

	users.Store["1"] = &User{
		Id:       1,
		Name:     "Leanne Graham",
		Username: "Bret",
		Email:    "Sincere@april.biz",
		Address: Address{
			Street:  "Kulas Light",
			Suite:   "Apt. 556",
			City:    "Gwenborough",
			Zipcode: "92998-3874",
			Geo: Geo{
				Lat: "-37.3159",
				Lng: "81.1496",
			},
		},
		Phone:   "1-770-736-8031 x56442",
		Website: "hildegard.org",
		Company: Company{
			Name:        "Romaguera-Crona",
			CatchPhrase: "Multi-layered client-server neural-net",
			Bs:          "harness real-time e-markets",
		},
		CreateDate: time.Date(2017, 3, 4, 05, 00, 00, 000, time.UTC),
	}

	users.Store["2"] = &User{
		Id:       2,
		Name:     "Ervin Howell",
		Username: "Antonette",
		Email:    "Shanna@melissa.tv",
		Address: Address{
			Street:  "Victor Plains",
			Suite:   "Suite 879",
			City:    "Wisokyburgh",
			Zipcode: "90566-7771",
			Geo: Geo{
				Lat: "-43.9509",
				Lng: "-34.4618",
			},
		},
		Phone:   "010-692-6593 x09125",
		Website: "anastasia.net",
		Company: Company{
			Name:        "Deckow-Crist",
			CatchPhrase: "Proactive didactic contingency",
			Bs:          "synergize scalable supply-chains",
		},
		CreateDate: time.Date(2017, 3, 4, 05, 00, 00, 000, time.UTC),
	}

	users.Store["3"] = &User{
		Id:       3,
		Name:     "Clementine Bauch",
		Username: "Samantha",
		Email:    "Nathan@yesenia.net",
		Address: Address{
			Street:  "Douglas Extension",
			Suite:   "Suite 847",
			City:    "McKenziehaven",
			Zipcode: "59590-4157",
			Geo: Geo{
				Lat: "-68.6102",
				Lng: "-47.0653",
			},
		},
		Phone:   "1-463-123-4447",
		Website: "ramiro.info",
		Company: Company{
			Name:        "Romaguera-Jacobson",
			CatchPhrase: "Face to face bifurcated interface",
			Bs:          "e-enable strategic applications",
		},
		CreateDate: time.Date(2017, 3, 4, 05, 00, 00, 000, time.UTC),
	}

	testApi := TestSimpleApi{
		Url:    "/users",
		Func:   users.GetUsers,
		Method: "GET",
		ReqUrl: "/users",
		Header: map[string]string{},
	}

	recorded := testApi.RunRequest(t)

	recorded.CodeIs(200)
	recorded.ContentTypeIsJson()
	recorded.BodyIs(`[
  {
    "id": 1,
    "name": "Leanne Graham",
    "username": "Bret",
    "email": "Sincere@april.biz",
    "address": {
      "street": "Kulas Light",
      "suite": "Apt. 556",
      "city": "Gwenborough",
      "zipcode": "92998-3874",
      "geo": {
        "lat": "-37.3159",
        "lng": "81.1496"
      }
    },
    "phone": "1-770-736-8031 x56442",
    "website": "hildegard.org",
    "company": {
      "name": "Romaguera-Crona",
      "catchPhrase": "Multi-layered client-server neural-net",
      "bs": "harness real-time e-markets"
    },
    "createDate": "2017-03-04T05:00:00Z"
  },
  {
    "id": 2,
    "name": "Ervin Howell",
    "username": "Antonette",
    "email": "Shanna@melissa.tv",
    "address": {
      "street": "Victor Plains",
      "suite": "Suite 879",
      "city": "Wisokyburgh",
      "zipcode": "90566-7771",
      "geo": {
        "lat": "-43.9509",
        "lng": "-34.4618"
      }
    },
    "phone": "010-692-6593 x09125",
    "website": "anastasia.net",
    "company": {
      "name": "Deckow-Crist",
      "catchPhrase": "Proactive didactic contingency",
      "bs": "synergize scalable supply-chains"
    },
    "createDate": "2017-03-04T05:00:00Z"
  },
  {
    "id": 3,
    "name": "Clementine Bauch",
    "username": "Samantha",
    "email": "Nathan@yesenia.net",
    "address": {
      "street": "Douglas Extension",
      "suite": "Suite 847",
      "city": "McKenziehaven",
      "zipcode": "59590-4157",
      "geo": {
        "lat": "-68.6102",
        "lng": "-47.0653"
      }
    },
    "phone": "1-463-123-4447",
    "website": "ramiro.info",
    "company": {
      "name": "Romaguera-Jacobson",
      "catchPhrase": "Face to face bifurcated interface",
      "bs": "e-enable strategic applications"
    },
    "createDate": "2017-03-04T05:00:00Z"
  }
]`)
}

func TestCallGetUserById(t *testing.T) {
	users := Users{
		Store: map[string]*User{},
	}

	users.Store["1"] = &User{
		Id:       1,
		Name:     "Leanne Graham",
		Username: "Bret",
		Email:    "Sincere@april.biz",
		Address: Address{
			Street:  "Kulas Light",
			Suite:   "Apt. 556",
			City:    "Gwenborough",
			Zipcode: "92998-3874",
			Geo: Geo{
				Lat: "-37.3159",
				Lng: "81.1496",
			},
		},
		Phone:   "1-770-736-8031 x56442",
		Website: "hildegard.org",
		Company: Company{
			Name:        "Romaguera-Crona",
			CatchPhrase: "Multi-layered client-server neural-net",
			Bs:          "harness real-time e-markets",
		},
		CreateDate: time.Date(2017, 3, 4, 05, 00, 00, 000, time.UTC),
	}

	testApi := TestSimpleApi{
		Url:    "/users/:id",
		Func:   users.GetUserById,
		Method: "GET",
		ReqUrl: "/users/1",
		Header: map[string]string{},
	}

	recorded := testApi.RunRequest(t)

	recorded.CodeIs(200)
	recorded.ContentTypeIsJson()
	recorded.BodyIs(`{
  "id": 1,
  "name": "Leanne Graham",
  "username": "Bret",
  "email": "Sincere@april.biz",
  "address": {
    "street": "Kulas Light",
    "suite": "Apt. 556",
    "city": "Gwenborough",
    "zipcode": "92998-3874",
    "geo": {
      "lat": "-37.3159",
      "lng": "81.1496"
    }
  },
  "phone": "1-770-736-8031 x56442",
  "website": "hildegard.org",
  "company": {
    "name": "Romaguera-Crona",
    "catchPhrase": "Multi-layered client-server neural-net",
    "bs": "harness real-time e-markets"
  },
  "createDate": "2017-03-04T05:00:00Z"
}`)
}
