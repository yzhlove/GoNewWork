package registry

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

const (
	ServerPort  = ":3000"
	ServicesURL = "http://localhost" + ServerPort + "/services"
)

type registry struct {
	registrations []Registration
	*sync.RWMutex
}

func (r registry) add(reg Registration) error {
	r.Lock()
	defer r.Unlock()

	r.registrations = append(r.registrations, reg)
	return nil
}

var _reg = registry{registrations: make([]Registration, 0, 4), RWMutex: new(sync.RWMutex)}

type Service struct{}

func (s Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received")
	switch r.Method {
	case http.MethodPost:
		decoder := json.NewDecoder(r.Body)
		var r Registration
		if err := decoder.Decode(&r); err != nil {
			log.Printf("decoder registration error, reason:%w", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("add service: %v with url: %s \n", r.ServiceName, r.ServiceURL)
		if err := _reg.add(r); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("add service succeed: %+v", r)))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
