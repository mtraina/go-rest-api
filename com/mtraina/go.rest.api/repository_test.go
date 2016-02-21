package main

import (
	//"testing"
	"github.com/boltdb/bolt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	//"time"
	"fmt"
	"testing"
	"log"
	"time"
)

var _ = Describe("Repository", func() {
	var dbName string = "my.db"
	var bucketName string = "todos"

	AfterEach(func() {
		log.Print("before")

		db, err := bolt.Open(dbName, 0600, &bolt.Options{Timeout: 1 * time.Second})
		Ω(err).NotTo(HaveOccurred())

		err = db.Update(func(tx *bolt.Tx) error {
			log.Print("delete bucket")

			err := tx.DeleteBucket([]byte(bucketName))

			log.Print("bucket deleted")

			log.Printf("err %s", err)

			Ω(err).NotTo(HaveOccurred())

			if err != nil {
				return fmt.Errorf("delete bucket: %s", err)
			}
			return nil
		})

		log.Print("after delete")
		Ω(err).NotTo(HaveOccurred())
	})

	Describe("Get from bucket", func() {
		Context("Get string", func() {

			It("should return a string", func() {
				//db, err := bolt.Open(dbName, 0600, &bolt.Options{Timeout: 1 * time.Second})

				//Ω(err).NotTo(HaveOccurred())

				//FindTodo()

				Ω(1).Should(Equal(1))

				//Ω(todo).To(BeEquivalentTo("Write presentationww"))
			})
		})
	})
})

func TestHandlers2(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handler suite")
}
