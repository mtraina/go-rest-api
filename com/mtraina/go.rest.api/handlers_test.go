package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"fmt"
	"encoding/json"
)

func TestTodoIndex(t *testing.T){
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Fatal("request creation error")
	}

	TodoIndex(w, r)

	if w.Code != 200 {
		t.Fatalf("wrong status code: {}", w.Code)
	}

	var ts Todos
	err = json.NewDecoder(w.Body).Decode(&ts)



	fmt.Printf("the returned body is {}", ts)
}

//func TestCreateTodo(t *testing.T){
//	TodoCreate()
//}
