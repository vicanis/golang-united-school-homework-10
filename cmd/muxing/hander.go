package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func createRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.HandleFunc("/name/{name}", func(w http.ResponseWriter, r *http.Request) {
		args := mux.Vars(r)

		if _, ok := args["name"]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Hello, %s!", args["name"])))
	}).Methods(http.MethodGet)

	router.HandleFunc("/bad", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}).Methods(http.MethodGet)

	router.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(append([]byte("I got message:\n"), data...))
	}).Methods(http.MethodPost)

	router.HandleFunc("/headers", func(w http.ResponseWriter, r *http.Request) {
		headers := r.Header

		if headers == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if _, ok := headers["A"]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if _, ok := headers["B"]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		headerA := headers["A"][0]
		headerB := headers["B"][0]

		intA, err := strconv.Atoi(headerA)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		intB, err := strconv.Atoi(headerB)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Add("a+b", fmt.Sprintf("%d", intA+intB))

		w.WriteHeader(http.StatusOK)
	}).Methods(http.MethodPost)

	return router
}
