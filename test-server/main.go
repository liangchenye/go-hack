package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	"sync"
)

func main() {
	init_db ()

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/os", GetAllOS),
		rest.Post("/os", PostOS),
		rest.Get("/os/:Distribution", GetOS),
		rest.Delete("/os/:Distribution", DeleteOS),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

type OS struct {
	ID string
	Distribution string
	Arch string
}

var store = map[string]*OS{}

var lock = sync.RWMutex{}

func GetOS(w rest.ResponseWriter, r *rest.Request) {
	Distribution := r.PathParam("Distribution")

	lock.RLock()
	var os *OS
	if store[Distribution] != nil {
		os = &OS{}
		*os = *store[Distribution]
	}
	lock.RUnlock()

	if os == nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(os)
}

func GetAllOS(w rest.ResponseWriter, r *rest.Request) {
	lock.RLock()
	os_list := make([]OS, len(store))
	i := 0
	for _, os := range store {
		os_list[i] = *os
		i++
	}
	lock.RUnlock()
	w.WriteJson(&os_list)
}

func PostOS(w rest.ResponseWriter, r *rest.Request) {
	os := OS{}
	err := r.DecodeJsonPayload(&os)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if os.Distribution == "" {
		rest.Error(w, "os Distribution required", 400)
		return
	}
	if os.Arch == "" {
		rest.Error(w, "os name required", 400)
		return
	}
	lock.Lock()
	store[os.Distribution] = &os
	lock.Unlock()
	w.WriteJson(&os)
}

func DeleteOS(w rest.ResponseWriter, r *rest.Request) {
	Distribution := r.PathParam("Distribution")
	lock.Lock()
	delete(store, Distribution)
	lock.Unlock()
	w.WriteHeader(http.StatusOK)
}

func init_db () {
	var os OS
	os.Distribution = "Ubuntu12.04"
	os.Arch = "x86_64"
	os.ID = "0001"
	store[os.Distribution] = &os
}
