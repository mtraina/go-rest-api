package main

import (
	"net/http"
	"fmt"
	"encoding/json"

	"testing"
	"net/http/httptest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"bytes"
)

var _ = Describe("Handlers", func() {

	var (
		w *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		w = httptest.NewRecorder()
	})

	Describe("Get all the todos", func() {
		Context("Get todos", func() {
			r, err := http.NewRequest("GET", "/", nil)

			It("should return 2 todos", func() {
				Ω(err).NotTo(HaveOccurred())

				TodoIndex(w, r)

				Ω(w.Code).To(BeEquivalentTo(http.StatusOK))

				var ts Todos
				err = json.NewDecoder(w.Body).Decode(&ts)

				Ω(len(ts)).To(BeEquivalentTo(2))

				fmt.Printf("the returned body is {}", ts)
			})
		})
	})

	Describe("Add a todo", func() {
		Context("Post a valid todo", func() {

			payload := []byte(`{"name":"Go on holiday", "completed":false}`)
			r, err := http.NewRequest("POST", "/todos", bytes.NewReader(payload))
			r.Header.Set("Content-Type", "application/json")

			It("should create a todo", func() {
				Ω(err).NotTo(HaveOccurred())

				TodoCreate(w, r)

				Ω(w.Code).To(BeEquivalentTo(http.StatusCreated))

				var t Todo
				err = json.NewDecoder(w.Body).Decode(&t)

				Ω(t.Id).To(BeEquivalentTo(3))
			})
		})

		Context("Post an invalid todo", func() {

			payload := []byte(`{"name":1, "completed":false}`)
			r, err := http.NewRequest("POST", "/todos", bytes.NewReader(payload))
			r.Header.Set("Content-Type", "application/json")

			It("should create a todo", func() {
				Ω(err).NotTo(HaveOccurred())

				TodoCreate(w, r)

				Ω(w.Code).To(BeEquivalentTo(422))
			})
		})
	})
})

func TestHandlers(t *testing.T){
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handler suite")
}