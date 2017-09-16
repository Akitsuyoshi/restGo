package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Index set a default endpoint for the root path
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome!")
}

//TodoIndex exports just string "Todo Index"
func TodoIndex(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

// TodoShow exports route, coresponding to the passed ID
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoID"]
	id, err := strconv.Atoi(todoID)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err = json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}
	t := RepoFindTodo(id)
	if t.ID == 0 && t.Name == "" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
	return
}

// TodoCreate create endpoint to the handlers file
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	// to limit request within 1MiB
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err = json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateTodo(todo)
	location := fmt.Sprintf("http://%s/%d", r.Host, t.ID)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

// TodoDelete return 204 status if Request is applicable, otherwise 422.
func TodoDelete(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["todoId"]
	id, err := strconv.Atoi(idParams)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		if err = json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	t := RepoFindTodo(id)
	if t.ID == 0 && t.Name == "" {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
	return
}
