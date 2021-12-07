package registry

import (
	"encoding/json"
	"fmt"
	"io"
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

func (r *registry) add(reg Registration) error {
	r.Lock()
	defer r.Unlock()

	r.registrations = append(r.registrations, reg)
	return nil
}

func (r *registry) remove(url string) error {
	for k, v := range r.registrations {
		if v.ServiceURL == url {
			r.Lock()
			r.registrations = append(r.registrations[:k], r.registrations[k+1:]...)
			r.Unlock()
			return nil
		}
	}
	return fmt.Errorf("service at url %s not found", url)
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

	case http.MethodDelete:
		payload, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("read request body error , reason:%w \n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		url := string(payload)
		log.Printf("removing service at utl:%s\n", url)

		if err := _reg.remove(url); err != nil {
			log.Printf("remove service error, reason:%v \n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
