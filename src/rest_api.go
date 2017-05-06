package main

import (
	"log"
	"net/http"
	"sync"
	"time"

	"fmt"

	"github.com/ant0ine/go-json-rest/rest"
)

type Geo struct {
	Lat string `bson:"lat" json:"lat"`
	Lng string `bson:"lng" json:"lng"`
}

type Address struct {
	Street  string `bson:"street" json:"street"`
	Suite   string `bson:"suite" json:"suite"`
	City    string `bson:"city" json:"city"`
	Zipcode string `bson:"zipcode" json:"zipcode"`
	Geo     Geo    `bson:"geo" json:"geo"`
}

type Company struct {
	Name        string `bson:"name" json:"name"`
	CatchPhrase string `bson:"catchPhrase" json:"catchPhrase"`
	Bs          string `bson:"bs" json:"bs"`
}

type User struct {
	Id         int       `bson:"id" json:"id"`
	Name       string    `bson:"name" json:"name"`
	Username   string    `bson:"username" json:"username"`
	Email      string    `bson:"email" json:"email"`
	Address    Address   `bson:"address" json:"address"`
	Phone      string    `bson:"phone" json:"phone"`
	Website    string    `bson:"website" json:"website"`
	Company    Company   `bson:"company" json:"company"`
	CreateDate time.Time `bson:"createDate" json:"createDate"`
}

type Users struct {
	sync.RWMutex
	Store map[string]*User
}

func index(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(map[string]string{"Test": "111"})
}

func main() {
	users := Users{
		Store: map[string]*User{},
	}

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/", index),
		rest.Get("/users", users.GetUsers),
		rest.Get("/users/:id", users.GetUserById),
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

func (u *Users) GetUserById(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	u.RLock()
	var user *User
	if u.Store[id] != nil {
		user = &User{}
		*user = *u.Store[id]
	}
	u.RUnlock()
	if user == nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(user)
}

func (u *Users) AddUsers(w rest.ResponseWriter, r *rest.Request) {
	user := User{}
	err := r.DecodeJsonPayload(&user)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u.Lock()
	id := len(u.Store)
	user.Id = id
	u.Store[fmt.Sprintf("%d", id)] = &user
	u.Unlock()
	w.WriteJson(&user)
}
