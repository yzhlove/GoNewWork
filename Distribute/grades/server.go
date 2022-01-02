package grades

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func RegistryHandlers() {
	handler := new(studentHandler)
	http.Handle("/students", handler)
	http.Handle("/students/", handler)
}

type studentHandler struct{}

// /students
// /students/{id}
// /student/{id}/grade
func (s studentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := strings.Split(r.URL.Path, "/")
	switch len(params) {
	case 2:
		s.All(w, r)
	case 3:
		id, err := strconv.Atoi(params[2])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		s.One(w, r, id)
	case 4:
		id, err := strconv.Atoi(params[2])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		s.Insert(w, r, id)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (s studentHandler) All(w http.ResponseWriter, r *http.Request) {
	_stuMutex.RLock()
	defer _stuMutex.RUnlock()

	data, err := s.toJSON(_students)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func (s studentHandler) One(w http.ResponseWriter, r *http.Request, id int) {
	_stuMutex.RLock()
	defer _stuMutex.RUnlock()

	stu, err := _students.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println(err)
		return
	}

	data, err := s.toJSON(stu)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func (s studentHandler) Insert(w http.ResponseWriter, r *http.Request, id int) {
	_stuMutex.Lock()
	defer _stuMutex.Unlock()

	stu, err := _students.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println(err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var grade Grade
	if err := decoder.Decode(&grade); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}

	stu.Grades = append(stu.Grades, grade)
	w.WriteHeader(http.StatusCreated)
	data, err := s.toJSON(stu)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func (s studentHandler) toJSON(input interface{}) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	encode := json.NewEncoder(buf)

	if err := encode.Encode(input); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
