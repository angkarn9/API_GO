package main

import (
	"log"
	"net/http"
	"sync"

	"fmt"

	"github.com/ant0ine/go-json-rest/rest"
)

type User struct {
	Id   string `bson:"id" json:"id"`
	Name string `bson:"name" json:"name"`
}

type Users struct {
	sync.RWMutex
	Store map[string]*User
}

func main() {
	users := Users{
		Store: map[string]*User{},
	}

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/users", users.GetUsers),
		rest.Post("/users", users.AddUsers),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":1323", api.MakeHandler()))
}

func (u *Users) GetUsers(w rest.ResponseWriter, r *rest.Request) {
	u.RLock()
	users := make([]User, len(u.Store))
	i := 0
	for _, user := range u.Store {
		users[i] = *user
		i++
	}
	u.RUnlock()
	w.WriteJson(&users)
}

func (u *Users) AddUsers(w rest.ResponseWriter, r *rest.Request) {
	user := User{}
	err := r.DecodeJsonPayload(&user)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u.Lock()
	id := fmt.Sprintf("%d", len(u.Store))
	user.Id = id
	u.Store[id] = &user
	u.Unlock()
	w.WriteJson(&user)
}
