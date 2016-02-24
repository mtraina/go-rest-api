package main

import (
	//"testing"
	"github.com/boltdb/bolt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	//"time"
	"testing"
	"log"
	"time"
)

var _ = Describe("Repository", func() {

	AfterEach(func() {
		log.Print("before")

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

		log.Print("after delete")
		Ω(err).NotTo(HaveOccurred())
	})

	Describe("Get from bucket", func() {
		Context("Get string", func() {

			It("should return a string", func() {
				//db, err := bolt.Open(dbName, 0600, &bolt.Options{Timeout: 1 * time.Second})

				//Ω(err).NotTo(HaveOccurred())

				todo := FindTodo()

				Ω(todo).Should(Equal("Write presentation"))

				//Ω(1).To(BeEquivalentTo(1))
			})
		})
	})
})

func TestHandlers2(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handler suite")
}
