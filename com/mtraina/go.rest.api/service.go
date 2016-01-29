package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"encoding/json"
)

type Todo struct {
	Name 		string	`json:"name"`
	Completed 	bool	`json:"completed"`
}

type Todos []Todo

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoId}", TodoShow)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	todos := Todos {
		Todo { Name: "Foo", Completed: false },
		Todo { Name: "Bar", Completed: true  },
	}

	json.NewEncoder(w).Encode(todos)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}