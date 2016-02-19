package main

import (
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/boltdb/bolt"
	"time"
)

var _ = Describe("Repository", func() {

	AfterEach(func() {

	})

	Describe("Get from bucket", func() {
		Context("Get string", func() {


			It("should return a string", func() {
				db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})

				Ω(err).NotTo(HaveOccurred())

			})
		})
	})
}


