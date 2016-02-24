package main

import (
	"github.com/boltdb/bolt"
	"log"
	"time"

	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Repository", func() {

	AfterEach(func() {
		log.Print("after test")

		db, err := bolt.Open(dbName, 0600, &bolt.Options{Timeout: 1 * time.Second})
		Ω(err).NotTo(HaveOccurred())

		err = db.Update(func(tx *bolt.Tx) error {
			log.Print("delete bucket")
			err := tx.DeleteBucket([]byte(bucketName))
			log.Print("bucket deleted")

			if err != nil {
				log.Printf("delete bucket error %v", err)
			}

			Ω(err).NotTo(HaveOccurred())
			return nil
		})

		log.Print("after delete bucket")
		Ω(err).NotTo(HaveOccurred())
	})

	Describe("Get from bucket", func() {
		Context("Get string", func() {

			It("should return a string", func() {
				todo := FindTodo()
				Ω(todo).Should(Equal("Write presentation"))
			})
		})
	})
})

func TestHandlers2(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handler suite")
}
