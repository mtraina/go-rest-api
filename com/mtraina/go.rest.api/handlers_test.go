package main

import (
	"net/http"
	"fmt"
	"encoding/json"

	"testing"
	"net/http/httptest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Handlers", func() {

	var (
		w *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		w = httptest.NewRecorder()
	})

	Describe("Get todos", func() {
		Context("Plain", func() {
			It("should return 2 todos", func() {
				r, err := http.NewRequest("GET", "/", nil)
				Ω(err).NotTo(HaveOccurred())

				TodoIndex(w, r)

				var ts Todos
				err = json.NewDecoder(w.Body).Decode(&ts)

				Ω(len(ts)).To(BeEquivalentTo(2))
				fmt.Printf("the returned body is {}", ts)
			})
		})
	})
})

func TestHandlers(t *testing.T){
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handler suite")
}

//func TestTodoIndex(t *testing.T){
//	w := httptest.NewRecorder()
//	r, err := http.NewRequest("GET", "/", nil)
//
//	if err != nil {
//		t.Fatal("request creation error")
//	}
//
//	TodoIndex(w, r)
//
//	if w.Code != 200 {
//		t.Fatalf("wrong status code: {}", w.Code)
//	}
//
//	var ts Todos
//	err = json.NewDecoder(w.Body).Decode(&ts)
//
//
//
//	fmt.Printf("the returned body is {}", ts)
//}

//func TestCreateTodo(t *testing.T){
//	TodoCreate()
//}
